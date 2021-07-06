package manager

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	svcpb "github.com/fregie/simple/proto/gen/go/simple-interface"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager struct {
	sessIDMap     sync.Map
	protoMap      map[string]*sync.Map
	svcMap        map[string]svcpb.InterfaceClient
	supportProtos []string
	db            *gorm.DB
	logger        *log.Logger
}

func NewManager(sqlitePath string) (*Manager, error) {
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Session{})
	if err != nil {
		return nil, err
	}
	m := &Manager{
		protoMap:      make(map[string]*sync.Map),
		svcMap:        make(map[string]svcpb.InterfaceClient),
		supportProtos: make([]string, 0),
		db:            db,
		logger:        log.Default(),
	}
	sessions := make([]Session, 0)
	err = db.Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	for i, sess := range sessions {
		m.sessIDMap.Store(sess.ID, &sessions[i])
		if _, ok := m.protoMap[sess.Proto]; !ok {
			m.protoMap[sess.Proto] = &sync.Map{}
		}
		m.protoMap[sess.Proto].Store(sess.Index, &sessions[i])
	}
	return m, nil
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
	protoMap, ok := m.protoMap[rsp.Name]
	if !ok {
		m.protoMap[rsp.Name] = &sync.Map{}
	}

	rsp2, err := svc.IsSupportPersistence(context.Background(), &svcpb.IsSupportPersistenceReq{})
	if err == nil && !rsp2.IsSupport {
		go m.syncProtoSessions(svc, protoMap, time.Minute)
	}

	return nil
}

func (m *Manager) GetProtos() []string {
	protos := make([]string, len(m.supportProtos))
	copy(protos, m.supportProtos)
	return protos
}

func (m *Manager) getService(name string) svcpb.InterfaceClient {
	svc, ok := m.svcMap[name]
	if !ok {
		return nil
	}
	return svc
}

func (m *Manager) CreateSession(ctx context.Context, proto string, configType svcpb.ConfigType, opt *svcpb.Option, customOpt string) ([]byte, error) {
	svc := m.getService(proto)
	if svc == nil {
		return nil, errors.New("Unknown proto")
	}
	rsp, err := svc.Create(ctx, &svcpb.CreateReq{
		ConfigType:   configType,
		Opt:          opt,
		CustomOption: customOpt,
	})
	if err != nil {
		return nil, fmt.Errorf("Create: %v", err)
	}
	if rsp.Code != svcpb.Code_OK {
		return nil, fmt.Errorf("Create %s", rsp.Msg)
	}

	sess := &Session{
		ID:           genSessionID(proto, rsp.Index),
		Proto:        proto,
		Index:        rsp.Index,
		ConfigType:   int32(rsp.Config.ConfigType),
		Config:       string(rsp.Config.Config),
		CustomOption: customOpt,
	}
	if opt != nil {
		sess.SendRateLimit = opt.SendRateLimit
		sess.RecvRateLimit = opt.RecvRateLimit
	}
	m.sessIDMap.Store(sess.ID, sess)
	m.protoMap[sess.Proto].Store(sess.Index, sess)
	err = m.db.Create(sess).Error
	if err != nil {
		m.logger.Printf("Create session to db failed: %s", err)
	}

	return rsp.Config.Config, nil
}

func (m *Manager) DeleteSession(ctx context.Context, sessID string) error {
	v, loaded := m.sessIDMap.LoadAndDelete(sessID)
	if !loaded {
		return fmt.Errorf("session [%s] not found", sessID)
	}
	sess := v.(*Session)
	m.protoMap[sess.Proto].Delete(sess.Index)
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
	err = m.db.Delete(sess).Error
	if err != nil {
		m.logger.Printf("Delete session to db failed: %s", err)
	}
	return nil
}

func (m *Manager) GetSession(sessID string) *Session {
	v, loaded := m.sessIDMap.Load(sessID)
	if !loaded {
		return nil
	}
	return v.(*Session)
}

func (m *Manager) GetAllSession() []*Session {
	sessions := make([]*Session, 0)
	m.sessIDMap.Range(func(k, v interface{}) bool {
		sess := v.(*Session)
		sessions = append(sessions, sess)
		return true
	})
	return sessions
}

func (m *Manager) syncProtoSessions(svc svcpb.InterfaceClient, protoMap *sync.Map, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), interval)
		protoMap.Range(func(k, v interface{}) bool {
			index := k.(string)
			rsp, err := svc.Get(ctx, &svcpb.GetReq{Index: index})
			if err != nil {
				m.logger.Print(err)
				return false
			}
			if rsp.Code != svcpb.Code_OK {
				sess := v.(*Session)
				rsp, err := svc.Create(ctx, &svcpb.CreateReq{
					ConfigType:   svcpb.ConfigType(sess.ConfigType),
					Opt:          sess.convertOption(),
					CustomOption: sess.CustomOption,
				})
				if err != nil {
					m.logger.Print(err)
					return false
				}
				if rsp.Code != svcpb.Code_OK {
					m.logger.Printf("sync: create: %s", rsp.Msg)
				}
			}
			return true
		})
		cancel()
	}
}
