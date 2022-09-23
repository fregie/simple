package manager

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	svcpb "github.com/fregie/simple/proto/gen/go/simple-interface"

	"github.com/cloudquery/sqlite"
	"gorm.io/gorm"
)

type Manager struct {
	host          string
	sessIDMap     sync.Map
	sessNameMap   sync.Map
	protoMap      map[string]*sync.Map
	svcMap        map[string]svcpb.InterfaceClient
	supportProtos []string
	db            *gorm.DB
	logger        *log.Logger
}

func NewManager(sqlitePath, host string) (*Manager, error) {
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Session{})
	if err != nil {
		return nil, err
	}
	m := &Manager{
		host:          host,
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
		m.sessNameMap.Store(sess.Name, &sessions[i])
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
		protoMap = &sync.Map{}
		m.protoMap[rsp.Name] = protoMap
	}
	rsp2, err := svc.IsSupportPersistence(context.Background(), &svcpb.IsSupportPersistenceReq{})
	if err == nil && !rsp2.IsSupport {
		go m.syncProtoSessions(svc, protoMap, time.Minute)
	}

	svc.SetMetadata(context.Background(), &svcpb.SetMetadataReq{Domain: m.host})

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

func (m *Manager) CreateSession(ctx context.Context, name, proto string, configType svcpb.ConfigType, opt *svcpb.Option, customOpt string) (*Session, error) {
	if name != "" {
		if _, ok := m.sessNameMap.Load(name); ok {
			return nil, fmt.Errorf("Create Session: session %s already exists", name)
		}
	}
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
		Name:         name,
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
	m.sessNameMap.Store(sess.Name, sess)
	m.protoMap[sess.Proto].Store(sess.Index, sess)
	err = m.db.Create(sess).Error
	if err != nil {
		m.logger.Printf("Create session to db failed: %s", err)
	}

	return sess, nil
}

func (m *Manager) DeleteSession(ctx context.Context, sessIDorName string) error {
	var sessID string
	v1, loaded := m.sessNameMap.Load(sessIDorName)
	if loaded {
		sess := v1.(*Session)
		sessID = sess.ID
	} else {
		sessID = sessIDorName
	}
	v, loaded := m.sessIDMap.LoadAndDelete(sessID)
	if !loaded {
		return fmt.Errorf("session [%s] not found", sessID)
	}
	sess := v.(*Session)
	m.sessNameMap.Delete(sess.Name)
	m.protoMap[sess.Proto].Delete(sess.Index)
	err := m.db.Delete(sess).Error
	if err != nil {
		m.logger.Printf("Delete session to db failed: %s", err)
	}
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

func (m *Manager) DeleteSessionByName(ctx context.Context, name string) error {
	v, loaded := m.sessNameMap.Load(name)
	if !loaded {
		return fmt.Errorf("session [%s] not found", name)
	}
	sess := v.(*Session)
	return m.DeleteSession(ctx, sess.ID)
}

func (m *Manager) GetSession(sessID string) *Session {
	v, loaded := m.sessIDMap.Load(sessID)
	if !loaded {
		return nil
	}
	return v.(*Session)
}

func (m *Manager) GetSessionByName(name string) *Session {
	v, loaded := m.sessNameMap.Load(name)
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

func (m *Manager) GetProtoSessions(proto string) ([]*Session, error) {
	sessMap, ok := m.protoMap[proto]
	if !ok {
		return nil, fmt.Errorf("Unknown proto %s", proto)
	}
	sessions := make([]*Session, 0)
	sessMap.Range(func(k, v interface{}) bool {
		sess := v.(*Session)
		sessions = append(sessions, sess)
		return true
	})
	return sessions, nil
}

func (m *Manager) syncProtoSessions(svc svcpb.InterfaceClient, protoMap *sync.Map, interval time.Duration) {
	timer := time.NewTimer(0)
	for range timer.C {
		ctx, cancel := context.WithTimeout(context.Background(), interval)
		svc.SetMetadata(ctx, &svcpb.SetMetadataReq{Domain: m.host})
		protoMap.Range(func(k, v interface{}) bool {
			index := k.(string)
			rsp, err := svc.Get(ctx, &svcpb.GetReq{Index: index})
			if err != nil {
				m.logger.Print(err)
				return false
			}
			if rsp.Code != svcpb.Code_OK {
				sess := v.(*Session)
				rsp, err := svc.CreateByConfig(ctx, &svcpb.CreateByConfigReq{
					Index: sess.Index,
					Opt: &svcpb.Option{
						SendRateLimit: sess.SendRateLimit,
						RecvRateLimit: sess.RecvRateLimit,
					},
					Config: &svcpb.Config{
						ConfigType: svcpb.ConfigType(sess.ConfigType),
						Config:     []byte(sess.Config),
					},
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
		timer.Reset(interval)
	}
}

func (m *Manager) GetSchemas(ctx context.Context) (map[string][]*svcpb.Field, error) {
	schemas := make(map[string][]*svcpb.Field)
	for proto, svc := range m.svcMap {
		rsp, err := svc.CustomOptionSchema(ctx, &svcpb.CustomOptionSchemaReq{})
		if err != nil {
			continue
		}
		schemas[proto] = rsp.Fields
	}
	return schemas, nil
}
