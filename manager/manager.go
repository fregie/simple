package manager

import (
	"context"
	"errors"
	"fmt"
	"sync"

	svcpb "github.com/fregie/simple-interface"
)

type Manager struct {
	sessMap       sync.Map
	svcMap        map[string]svcpb.InterfaceClient
	supportProtos []string
}

func NewManager() *Manager {
	return &Manager{
		svcMap:        make(map[string]svcpb.InterfaceClient),
		supportProtos: make([]string, 0),
	}
}

func (m *Manager) RegisterService(svc svcpb.InterfaceClient) error {
	if svc == nil {
		return errors.New("Register Service: service is nil")
	}
	rsp, err := svc.Name(context.Background(), &svcpb.NameReq{})
	if err != nil {
		return fmt.Errorf("Register Service: %v", err)
	}
	if rsp.Name == "" {
		return errors.New("Register Service: service name is empty")
	}
	m.svcMap[rsp.Name] = svc
	m.supportProtos = append(m.supportProtos, rsp.Name)

	return nil
}

func (m *Manager) getService(name string) svcpb.InterfaceClient {
	svc, ok := m.svcMap[name]
	if !ok {
		return nil
	}
	return svc
}

func (m *Manager) CreateSession(ctx context.Context, proto string, configType svcpb.ConfigType, opt svcpb.Option, customOpt string) ([]byte, error) {
	svc := m.getService(proto)
	if svc == nil {
		return nil, errors.New("Unknown proto")
	}
	rsp, err := svc.Create(ctx, &svcpb.CreateReq{
		ConfigType:   configType,
		Opt:          &opt,
		CustomOption: customOpt,
	})
	if err != nil {
		return nil, fmt.Errorf("Create: %v", err)
	}
	if rsp.Code != svcpb.Code_OK {
		return nil, fmt.Errorf("Create %s", rsp.Msg)
	}

	sess := &Session{
		ID:            genSessionID(proto, rsp.Index),
		Proto:         proto,
		Index:         rsp.Index,
		ConfigType:    int32(configType),
		SendRateLimit: opt.SendRateLimit,
		RecvRateLimit: opt.RecvRateLimit,
		CustomOption:  customOpt,
	}
	m.sessMap.Store(sess.ID, sess)

	return rsp.Config.Config, nil
}

func (m *Manager) DeleteSession(ctx context.Context, sessID string) error {
	v, loaded := m.sessMap.LoadAndDelete(sessID)
	if !loaded {
		return fmt.Errorf("session [%s] not found", sessID)
	}
	sess := v.(*Session)
	svc := m.getService(sess.Proto)
	if svc == nil {
		return errors.New("Unknown proto")
	}
	rsp, err := svc.Delete(ctx, &svcpb.DeleteReq{Index: sess.Index})
	if err != nil {
		return err
	}
	if rsp.Code != svcpb.Code_OK {
		return fmt.Errorf("Delete %s", rsp.Msg)
	}
	return nil
}

func (m *Manager) GetSession(sessID string) *Session {
	v, loaded := m.sessMap.Load(sessID)
	if !loaded {
		return nil
	}
	return v.(*Session)
}

func (m *Manager) GetAllSession() []*Session {
	sessions := make([]*Session, 0)
	m.sessMap.Range(func(k, v interface{}) bool {
		sess := v.(*Session)
		sessions = append(sessions, sess)
		return true
	})
	return sessions
}
