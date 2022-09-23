package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	tool "github.com/fregie/gotool"
	"github.com/fregie/simple/adapter/shadowsocks/ss"
	svcpb "github.com/fregie/simple/proto/simple-interface"
)

const Name = "ss"

type Service struct {
	svcpb.UnimplementedInterfaceServer
	host    string
	manager *ss.Manager
}

type CustomOption struct {
	Port     int    `json:"port"`
	Password string `json:"password"`
	Method   string `json:"method"`
}

// NewService returns a new shadowsocks adapter service.
func NewService(minPort, maxPort int) (*Service, error) {
	m := ss.NewManager(minPort, maxPort)
	return &Service{
		manager: m,
	}, nil
}

// Name returns the name of the adapter.
func (s *Service) Name(_ context.Context, _ *svcpb.NameReq) (*svcpb.NameRsp, error) {
	return &svcpb.NameRsp{Name: Name}, nil
}

// IsSupportPersistent returns true if the adapter supports persistent connections.
func (s *Service) IsSupportPersistence(_ context.Context, _ *svcpb.IsSupportPersistenceReq) (*svcpb.IsSupportPersistenceRsp, error) {
	return &svcpb.IsSupportPersistenceRsp{IsSupport: false}, nil
}

// CustomOptionSchema returns the schema of the custom options.
func (s *Service) CustomOptionSchema(_ context.Context, _ *svcpb.CustomOptionSchemaReq) (*svcpb.CustomOptionSchemaRsp, error) {
	rsp := &svcpb.CustomOptionSchemaRsp{
		Fields: []*svcpb.Field{
			{Name: "port", Type: svcpb.Type_Number},
			{Name: "password", Type: svcpb.Type_String},
			{Name: "method", Type: svcpb.Type_String},
		},
	}
	return rsp, nil
}

// SetMetadata sets the metadata of the adapter.
func (s *Service) SetMetadata(_ context.Context, req *svcpb.SetMetadataReq) (*svcpb.SetMetadataRsp, error) {
	if req.Domain != "" {
		s.manager.SetHost(req.Domain)
	} else if req.IP != "" {
		s.manager.SetHost(req.IP)
	}
	return &svcpb.SetMetadataRsp{}, nil
}

// Create creates a new ss server.
func (s *Service) Create(_ context.Context, req *svcpb.CreateReq) (rsp *svcpb.CreateRsp, e error) {
	rsp = &svcpb.CreateRsp{Code: svcpb.Code_OK}
	copt := CustomOption{
		Port:     0,
		Method:   "aes-128-gcm",
		Password: string(tool.RandomString(16)),
	}
	if req.CustomOption != "" {
		err := json.Unmarshal([]byte(req.CustomOption), &copt)
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = err.Error()
			return
		}
	}
	var limitSend, limitRecv uint64
	if req.Opt != nil {
		limitSend, limitRecv = req.Opt.SendRateLimit, req.Opt.RecvRateLimit
	}
	ss, err := s.manager.Add(copt.Port, copt.Method, copt.Password, ss.WithRateLimit(limitSend, limitRecv))
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}

	rsp.Index = strconv.Itoa(ss.Port)
	rsp.Config = &svcpb.Config{}
	switch req.ConfigType {
	case svcpb.ConfigType_JSON:
		rsp.Config.ConfigType = svcpb.ConfigType_JSON
		rsp.Config.Config = ss.ExportJson()
	case svcpb.ConfigType_URL:
		rsp.Config.ConfigType = svcpb.ConfigType_URL
		rsp.Config.Config = []byte(ss.ExportURL())
	default:
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = "unknown config type"
	}

	return
}

// CreateByConfig creates a new ss server by config.
func (s *Service) CreateByConfig(_ context.Context, req *svcpb.CreateByConfigReq) (rsp *svcpb.CreateByConfigRsp, e error) {
	rsp = &svcpb.CreateByConfigRsp{Code: svcpb.Code_OK}
	copt := CustomOption{}
	if req.CustomOption != "" {
		err := json.Unmarshal([]byte(req.CustomOption), &copt)
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = err.Error()
			return
		}
	}
	var conf ss.SS
	switch req.Config.ConfigType {
	case svcpb.ConfigType_JSON:
		err := json.Unmarshal(req.Config.Config, &conf)
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = fmt.Sprintf("parse json config: %s", err.Error())
			return
		}
	case svcpb.ConfigType_URL:
		var str string
		log.Print(string(req.Config.Config))
		_, err := fmt.Sscanf(string(req.Config.Config), "ss://%s", &str)
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = fmt.Sprintf("parse url config: %s", err.Error())
			return
		}
		decoded := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
		n, err := base64.StdEncoding.Decode(decoded, []byte(str))
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = fmt.Sprintf("decode base64 string: %s", err.Error())
			return
		}
		log.Print(string(decoded[:n]))
		parsed, err := url.Parse("ss://" + string(decoded[:n]))
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = fmt.Sprintf("parse url config: %s", err.Error())
			return
		}
		conf.Method = parsed.User.Username()
		conf.Password, _ = parsed.User.Password()
		conf.Server = parsed.Host
		conf.Port, err = strconv.Atoi(parsed.Port())
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = fmt.Sprintf("parse url config: %s", err.Error())
			return
		}

	default:
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = "unknown config type"
		return
	}
	_, err := s.manager.Add(conf.Port, conf.Method, conf.Password)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	return
}

// Get gets the ss server.
func (s *Service) Get(_ context.Context, req *svcpb.GetReq) (rsp *svcpb.GetRsp, e error) {
	rsp = &svcpb.GetRsp{Code: svcpb.Code_OK}
	port, err := strconv.Atoi(req.Index)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	ss, err := s.manager.Get(port)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	rsp.Config = &svcpb.Config{}
	switch req.ConfigType {
	case svcpb.ConfigType_JSON:
		rsp.Config.ConfigType = svcpb.ConfigType_JSON
		rsp.Config.Config = ss.ExportJson()
	case svcpb.ConfigType_URL:
		rsp.Config.ConfigType = svcpb.ConfigType_URL
		rsp.Config.Config = []byte(ss.ExportURL())
	default:
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = "unknown config type"
	}
	return
}

// Delete deletes the ss server.
func (s *Service) Delete(_ context.Context, req *svcpb.DeleteReq) (rsp *svcpb.DeleteRsp, e error) {
	rsp = &svcpb.DeleteRsp{Code: svcpb.Code_OK}
	port, err := strconv.Atoi(req.Index)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	err = s.manager.Del(port)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	return
}

// GetAll gets all ss servers.
func (s *Service) GetAll(_ context.Context, req *svcpb.GetAllReq) (rsp *svcpb.GetAllRsp, e error) {
	rsp = &svcpb.GetAllRsp{Code: svcpb.Code_OK}
	sss := s.manager.GetAll()
	rsp.All = make(map[string]*svcpb.Config)
	for _, ss := range sss {
		rsp.All[strconv.Itoa(ss.Port)] = &svcpb.Config{}
		switch req.ConfigType {
		case svcpb.ConfigType_JSON:
			rsp.All[strconv.Itoa(ss.Port)].ConfigType = svcpb.ConfigType_JSON
			rsp.All[strconv.Itoa(ss.Port)].Config = ss.ExportJson()
		case svcpb.ConfigType_URL:
			rsp.All[strconv.Itoa(ss.Port)].ConfigType = svcpb.ConfigType_URL
			rsp.All[strconv.Itoa(ss.Port)].Config = []byte(ss.ExportURL())
		default:
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = "unknown config type"
		}
	}
	return
}

// UpdateOption updates the ss server option.
func (s *Service) UpdateOption(_ context.Context, req *svcpb.UpdateOptionReq) (rsp *svcpb.UpdateOptionRsp, e error) {
	rsp = &svcpb.UpdateOptionRsp{Code: svcpb.Code_OK}
	port, err := strconv.Atoi(req.Index)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	ss, err := s.manager.Get(port)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	ss.SetRatelimit(float64(req.Opt.SendRateLimit), float64(req.Opt.RecvRateLimit),
		int64(req.Opt.SendRateLimit)*2, int64(req.Opt.RecvRateLimit)*2)
	return
}
