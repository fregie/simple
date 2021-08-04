package ss

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/fregie/freconn"
	"github.com/shadowsocks/go-shadowsocks2/socks"
)

// Packet NAT table
type natmap struct {
	sync.RWMutex
	m       map[string]net.PacketConn
	timeout time.Duration
}

func newNATmap(timeout time.Duration) *natmap {
	m := &natmap{}
	m.m = make(map[string]net.PacketConn)
	m.timeout = timeout
	return m
}

func (m *natmap) Get(key string) net.PacketConn {
	m.RLock()
	defer m.RUnlock()
	return m.m[key]
}

func (m *natmap) Set(key string, pc net.PacketConn) {
	m.Lock()
	defer m.Unlock()

	m.m[key] = pc
}

func (m *natmap) Del(key string) net.PacketConn {
	m.Lock()
	defer m.Unlock()

	pc, ok := m.m[key]
	if ok {
		delete(m.m, key)
		return pc
	}
	return nil
}

func (m *natmap) Add(peer net.Addr, dst, src net.PacketConn, role mode) {
	m.Set(peer.String(), src)

	go func() {
		timedCopy(dst, peer, src, m.timeout, role)
		if pc := m.Del(peer.String()); pc != nil {
			pc.Close()
		}
	}()
}

type mode int

const (
	remoteServer mode = iota
	relayClient
	socksClient
)

func timedCopy(dst net.PacketConn, target net.Addr, src net.PacketConn, timeout time.Duration, role mode) error {
	buf := make([]byte, udpBufSize)

	for {
		src.SetReadDeadline(time.Now().Add(timeout))
		n, raddr, err := src.ReadFrom(buf)
		if err != nil {
			return err
		}

		switch role {
		case remoteServer: // server -> client: add original packet source
			srcAddr := socks.ParseAddr(raddr.String())
			copy(buf[len(srcAddr):], buf[:n])
			copy(buf, srcAddr)
			if udpBufSize < len(srcAddr)+n {
				continue
			}
			_, err = dst.WriteTo(buf[:len(srcAddr)+n], target)
		case relayClient: // client -> user: strip original packet source
			srcAddr := socks.SplitAddr(buf[:n])
			_, err = dst.WriteTo(buf[len(srcAddr):n], target)
		case socksClient: // client -> socks5 program: just set RSV and FRAG = 0
			_, err = dst.WriteTo(append([]byte{0, 0, 0}, buf[:n]...), target)
		}

		if err != nil {
			return err
		}
	}
}

func handlePacketConn(ctx context.Context, ciphPc net.PacketConn, bindAddr string, opts ...OptionHandler) {
	defer ciphPc.Close()
	opt := &Option{}
	for _, opth := range opts {
		opth(opt)
	}
	t, _ := time.ParseDuration(defaultTimeout)
	nm := newNATmap(t)

	// Debug.Printf("listening packetConn on %s", ciphPc.LocalAddr().String())
	pcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			ciphPc.Close()
		case <-pcCtx.Done():
		}
	}()
	buf := make([]byte, udpBufSize)
	for {
		select {
		case <-ctx.Done():
			Debug.Printf("Stop packetConn on [%s] success", ciphPc.LocalAddr().String())
			return
		default:
			n, raddr, err := ciphPc.ReadFrom(buf)
			if err != nil {
				Debug.Printf("PacketConn remote read error: %v", err)
				continue
			}
			tgtAddr := socks.SplitAddr(buf[:n])
			if tgtAddr == nil {
				Debug.Printf("failed to split target address from packet: %q", buf[:n])
				continue
			}
			tgtUDPAddr, err := net.ResolveUDPAddr("udp", tgtAddr.String())
			if err != nil {
				Debug.Printf("failed to resolve target UDP address: %v", err)
				continue
			}
			payload := buf[len(tgtAddr):n]
			pc := nm.Get(raddr.String())
			if pc == nil {
				if bindAddr == "" {
					pc, err = net.ListenPacket("udp", "")
				} else {
					pc, err = net.ListenPacket("udp", bindAddr+":0")
				}
				if err != nil {
					Debug.Printf("UDP remote listen error: %v", err)
					continue
				}

				var rpc net.PacketConn
				if opt.EnableTrafficControl {
					rpc = freconn.WrapPacketConn(pc, freconn.WithLimit(opt.TxBucket, opt.RxBucket))
				} else {
					rpc = pc
				}
				nm.Add(raddr, ciphPc, rpc, remoteServer)
			}
			_, err = pc.WriteTo(payload, tgtUDPAddr)
			if err != nil {
				Debug.Printf("UDP remote write error: %v", err)
				continue
			}
		}
	}
}
