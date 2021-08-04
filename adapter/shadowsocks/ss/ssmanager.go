package ss

import (
	"fmt"
	"net"
	"sync"
)

type Manager struct {
	host    string
	minPort int
	maxPort int
	ssMap   sync.Map
}

func NewManager(minPort, maxPort int) *Manager {
	return &Manager{
		minPort: minPort,
		maxPort: maxPort,
	}
}

// SetHost set host
func (m *Manager) SetHost(host string) {
	m.host = host
}

func (m *Manager) Add(port int, method, password string) (*SS, error) {
	if m.host == "" {
		return nil, fmt.Errorf("host not set, please SetHost first")
	}
	if port == 0 {
		port = m.findAvailPort()
	}
	if port < m.minPort || port > m.maxPort {
		return nil, fmt.Errorf("port %d out of range", port)
	}
	ss := NewSS(m.host, port, method, password, &net.Dialer{})
	if _, exist := m.ssMap.LoadOrStore(port, ss); exist {
		return nil, fmt.Errorf("port %d already in use", port)
	}
	err := ss.RunTCP()
	if err != nil {
		return nil, err
	}
	err = ss.RunUDP()
	if err != nil {
		return nil, err
	}
	Info.Printf("Start shadowsocks [%s|%s|%d]", method, password, port)
	return ss, nil
}

func (m *Manager) Get(port int) (*SS, error) {
	ss, exist := m.ssMap.Load(port)
	if !exist {
		return nil, fmt.Errorf("port %d not exist", port)
	}
	return ss.(*SS), nil
}

// Del  delete ss by port
func (m *Manager) Del(port int) error {
	ss, exist := m.ssMap.LoadAndDelete(port)
	if !exist {
		return fmt.Errorf("port %d not exist", port)
	}
	ss.(*SS).StopTCP()
	ss.(*SS).StopUDP()
	return nil
}

// GetAll get all ss
func (m *Manager) GetAll() []*SS {
	ssSlice := make([]*SS, 0)
	m.ssMap.Range(func(key, value interface{}) bool {
		ssSlice = append(ssSlice, value.(*SS))
		return true
	})
	return ssSlice
}

// GetAllByMethod get all ss by method
func (m *Manager) GetAllByMethod(method string) []*SS {
	ssSlice := make([]*SS, 0)
	m.ssMap.Range(func(key, value interface{}) bool {
		ss := value.(*SS)
		if ss.Method == method {
			ssSlice = append(ssSlice, ss)
		}
		return true
	})
	return ssSlice
}

// findAvailPort find avail port
func (m *Manager) findAvailPort() int {
	for i := m.minPort; i <= m.maxPort; i++ {
		if _, exist := m.ssMap.Load(i); !exist {
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", i))
			if err != nil {
				continue
			}
			lis.Close()
			laddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%d", i))
			pc, err := net.ListenUDP("udp", laddr)
			if err != nil {
				continue
			}
			pc.Close()
			return i
		}
	}
	return 0
}
