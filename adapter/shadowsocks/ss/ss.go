package ss

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/shadowsocks/go-shadowsocks2/core"
	"golang.org/x/time/rate"
)

const (
	udpBufSize     = 4 * 1024
	defaultTimeout = "60s"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

type Dialer interface {
	Dial(network, addr string) (c net.Conn, err error)
}

func SetLog(debug, info, err *log.Logger) {
	Debug = debug
	Info = info
	Error = err
}

var logger = log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags)

// SS is a sadowsocks instance manager
type SS struct {
	Server         string `json:"server"`
	Port           int    `json:"server_port"`
	Method         string `json:"method"`
	Password       string `json:"password"`
	tcpRunning     bool
	udpCtx         context.Context
	udpCancel      context.CancelFunc
	listener       net.Listener
	udpPackegeConn net.PacketConn
	isRatelimit    bool
	txLimit        *rate.Limiter
	rxLimit        *rate.Limiter
	dialer         Dialer
}

// relay copies between left and right bidirectionally. Returns number of
// bytes copied from right to left, from left to right, and any error occurred.
func relay(left, right net.Conn) (int64, int64, error) {
	type res struct {
		N   int64
		Err error
	}
	ch := make(chan res)

	go func() {
		var n int64
		var err error
		// n, err = copyAndPeep(right, left, cp.ClientWrite)
		n, err = io.Copy(right, left)
		right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
		left.SetDeadline(time.Now())  // wake up the other goroutine blocking on left
		ch <- res{n, err}
	}()
	var n int64
	var err error
	// n, err = copyAndPeep(left, right, cp.ServerWrite)
	n, err = io.Copy(left, right)
	right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
	left.SetDeadline(time.Now())  // wake up the other goroutine blocking on left
	rs := <-ch

	if err == nil {
		err = rs.Err
	}
	return n, rs.N, err
}

// NewSS return a new SS pointer
func NewSS(server string, port int, method, password string, dialer Dialer) *SS {
	udpCtx, udpCancel := context.WithCancel(context.Background())
	return &SS{
		Server:      server,
		Port:        port,
		Method:      method,
		Password:    password,
		isRatelimit: false,
		udpCtx:      udpCtx,
		udpCancel:   udpCancel,
		dialer:      dialer,
	}
}

func (ss *SS) SetRatelimit(tRate, rRate float64, tBurst, rBurst int64) {
	ss.txLimit = rate.NewLimiter(rate.Limit(tRate), int(tBurst))
	ss.rxLimit = rate.NewLimiter(rate.Limit(rRate), int(rBurst))
	ss.isRatelimit = true
}

// RunTCP run ss server on tcp, use localAddr as source ip
func (ss *SS) RunTCP() error {
	var err error
	ciph, err := core.PickCipher(ss.Method, []byte{}, ss.Password)
	if err != nil {
		Error.Printf("Create SS on port [%d] failed: %s", ss.Port, err)
		return err
	}
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", ss.Port)

	ss.listener, err = net.Listen("tcp", addr)
	if err != nil {
		Error.Printf("Create SS on port [%d] failed: %s", ss.Port, err)
		Debug.Printf("failed to listen on %s : %v", addr, err)
		return err
	}
	Debug.Printf("Create SS on port [%d] success!", ss.Port)
	Debug.Printf("listening TCP on %s", addr)
	ss.tcpRunning = true
	go func() {
		for {
			c, err := ss.listener.Accept()
			if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					Debug.Printf("Stop SS on port [TCP:%d] success!", ss.Port)
					ss.tcpRunning = false
					return
				}
				Debug.Printf("failed to accept: %v", err)
				continue
			}
			c.(*net.TCPConn).SetKeepAlive(true)
			go func() {
				c = ciph.StreamConn(c)
				if ss.isRatelimit {
					err = handleConn(c, withTrafficControl(ss.txLimit, ss.rxLimit), withDialer(ss.dialer))
				} else {
					err = handleConn(c, withDialer(ss.dialer))
				}
			}()
		}
	}()

	return nil
}

// StopTCP stop ss server on tcp
func (ss *SS) StopTCP() error {
	if !ss.tcpRunning {
		Debug.Printf("SS[%d] is not running", ss.Port)
		return nil
	}
	Debug.Printf("SS is running, stopping on [%d]", ss.Port)
	return ss.listener.Close()
}

// RunUDP Listen on addr for encrypted packets and basically do UDP NAT.
func (ss *SS) RunUDP() error {
	var err error
	ciph, err := core.PickCipher(ss.Method, []byte{}, ss.Password)
	if err != nil {
		Error.Printf("Create SS on port [%d] failed: %s", ss.Port, err)
		return err
	}
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", ss.Port)
	ss.udpPackegeConn, err = net.ListenPacket("udp", addr)
	if err != nil {
		Debug.Printf("UDP remote listen error: %v", err)
		return err
	}
	ss.udpPackegeConn = ciph.PacketConn(ss.udpPackegeConn)
	localAddr := "0.0.0.0"
	if ss.isRatelimit {
		go handlePacketConn(ss.udpCtx, ss.udpPackegeConn, localAddr, withTrafficControl(ss.txLimit, ss.rxLimit))
	} else {
		go handlePacketConn(ss.udpCtx, ss.udpPackegeConn, localAddr)
	}

	return nil
}

func (ss *SS) StopUDP() error {
	Debug.Printf("SS is running, stopping on [UDP:%d]", ss.Port)
	ss.udpCancel()
	return nil
}

// ExportJson export ss server info to json
func (ss *SS) ExportJson() []byte {
	r, _ := json.Marshal(ss)
	return r
}

// ExportURL export ss server info to url
func (ss *SS) ExportURL() string {
	str := fmt.Sprintf("%s:%s@%s:%d", ss.Method, ss.Password, ss.Server, ss.Port)
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(encoded, []byte(str))
	return "ss://" + string(encoded)
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func init() {
	if Debug == nil {
		Debug = log.New(os.Stdout, "[Debug] ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	if Info == nil {
		Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	}
	if Error == nil {
		Error = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
