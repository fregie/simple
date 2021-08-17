package ss

import (
	"net"

	"github.com/fregie/freconn"
	"github.com/shadowsocks/go-shadowsocks2/socks"
	"golang.org/x/time/rate"
)

type SSOptionHandler func(*SSOption)
type SSOption struct {
	EnableBlocklist      bool
	Blocklist            map[string]bool
	EnableTrafficControl bool
	TxBucket             *rate.Limiter
	RxBucket             *rate.Limiter
	EnablePeeper         bool
	Dialer               Dialer
}

func withBlocklist(blocklist map[string]bool) SSOptionHandler {
	return func(opt *SSOption) {
		opt.EnableBlocklist = true
		opt.Blocklist = blocklist
	}
}

func withTrafficControl(txBucket, rxBucket *rate.Limiter) SSOptionHandler {
	return func(opt *SSOption) {
		opt.EnableTrafficControl = true
		opt.TxBucket = txBucket
		opt.RxBucket = rxBucket
	}
}

func withDialer(dialer Dialer) SSOptionHandler {
	return func(opt *SSOption) {
		if dialer != nil {
			opt.Dialer = dialer
		} else {
			opt.Dialer = &net.Dialer{}
		}
	}
}

func handleConn(ciphConn net.Conn, opts ...SSOptionHandler) error {
	defer ciphConn.Close()
	opt := &SSOption{}
	for _, opth := range opts {
		opth(opt)
	}
	tgt, err := socks.ReadAddr(ciphConn)
	if err != nil {
		return err
	}
	domain, _, _ := net.SplitHostPort(tgt.String())
	if opt.EnableBlocklist {
		if forbidden, ok := opt.Blocklist[domain]; ok && forbidden {
			// ciphConn.Write([]byte("Forbidden target"))
			return nil
		}
	}
	var rc net.Conn
	if opt.Dialer == nil {
		opt.Dialer = &net.Dialer{}
	}
	rc, err = opt.Dialer.Dial("tcp", tgt.String())
	if err != nil {
		Debug.Printf("failed to connect to target[%s]: %v", tgt.String(), err)
		return err
	}

	defer rc.Close()
	// rc.SetKeepAlive(true)

	var remoteConn net.Conn
	if opt.EnableTrafficControl && opt.RxBucket != nil && opt.TxBucket != nil {
		remoteConn = freconn.WrapConn(rc, freconn.WithLimit(opt.RxBucket, opt.TxBucket))
	} else {
		remoteConn = rc
	}
	_, _, err = relay(ciphConn, remoteConn)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil // ignore i/o timeout
		}
		Debug.Printf("relay error: %v", err)
		return err
	}
	return nil
}
