package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	svcpb "github.com/fregie/simple/proto/simple-interface"

	tool "github.com/fregie/gotool"
	trojanpb "github.com/p4gefau1t/trojan-go/api/service"
	"github.com/p4gefau1t/trojan-go/common"
	"google.golang.org/grpc"
)

const Name = "trojan"

type Service struct {
	srv  trojanpb.TrojanServerServiceClient
	conf TrojanServerConfig
	host string
	svcpb.UnimplementedInterfaceServer
}

func init() {
	rand.Seed(time.Now().Unix())
}

func NewService(configPath, addr string) (*Service, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("open trojan config:" + err.Error())
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read trojan config:" + err.Error())
	}
	manager := &Service{}
	err = json.Unmarshal(data, &manager.conf)
	if err != nil {
		return nil, fmt.Errorf("parse trajan json " + err.Error())
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	manager.srv = trojanpb.NewTrojanServerServiceClient(conn)
	return manager, nil
}

func (s *Service) Name(_ context.Context, _ *svcpb.NameReq) (*svcpb.NameRsp, error) {
	return &svcpb.NameRsp{Name: Name}, nil
}

func (s *Service) IsSupportPersistence(_ context.Context, _ *svcpb.IsSupportPersistenceReq) (*svcpb.IsSupportPersistenceRsp, error) {
	return &svcpb.IsSupportPersistenceRsp{IsSupport: s.conf.Sqlite != ""}, nil
}

func (s *Service) CustomOptionSchema(_ context.Context, _ *svcpb.CustomOptionSchemaReq) (*svcpb.CustomOptionSchemaRsp, error) {
	rsp := &svcpb.CustomOptionSchemaRsp{
		Fields: []*svcpb.Field{
			{Name: "ssl_verify", Type: svcpb.Type_Bool},
			{Name: "ssl_sni", Type: svcpb.Type_String},
		},
	}
	if s.conf.Websocket.Enable {
		rsp.Fields = append(rsp.Fields, &svcpb.Field{Name: "enable_websocket", Type: svcpb.Type_Bool})
		rsp.Fields = append(rsp.Fields, &svcpb.Field{Name: "websocket_path", Type: svcpb.Type_String})
		rsp.Fields = append(rsp.Fields, &svcpb.Field{Name: "websocket_host", Type: svcpb.Type_String})
	}
	if s.conf.Mux.Enable {
		rsp.Fields = append(rsp.Fields, &svcpb.Field{Name: "enable_mux", Type: svcpb.Type_Bool})
		rsp.Fields = append(rsp.Fields, &svcpb.Field{Name: "mux_concurrency", Type: svcpb.Type_Number})
		rsp.Fields = append(rsp.Fields, &svcpb.Field{Name: "mux_idle_timeout", Type: svcpb.Type_Number})
	}
	return rsp, nil
}

func (s *Service) SetMetadata(_ context.Context, req *svcpb.SetMetadataReq) (rsp *svcpb.SetMetadataRsp, e error) {
	if req.Domain != "" {
		s.host = req.Domain
	} else {
		s.host = req.IP
	}
	return &svcpb.SetMetadataRsp{}, nil
}

func (s *Service) Create(ctx context.Context, req *svcpb.CreateReq) (rsp *svcpb.CreateRsp, e error) {
	rsp = &svcpb.CreateRsp{Code: svcpb.Code_OK}
	var copt CustomOption
	if req.CustomOption != "" {
		err := json.Unmarshal([]byte(req.CustomOption), &copt)
		if err != nil {
			rsp.Code = svcpb.Code_Fail
			rsp.Msg = err.Error()
			return
		}
	}
	c := &TrojanClientConfig{
		RunType:    "client",
		RemoteAddr: s.host,
		RemotePort: s.conf.LocalPort,
		Password:   []string{string(tool.RandomString(16))},
		SSL: SSL{
			Verify: copt.SslVerify,
			SNI:    copt.SslSni,
		},
		Mux: Mux{
			Enable:      copt.EnableMux,
			Concurrency: copt.MuxConCurrency,
			IdleTimeout: copt.MuxIdleTimeout,
		},
		Websocket: Websocket{
			Enable: copt.EnableWS,
			Path:   copt.WSPath,
			Host:   copt.WSHost,
		},
	}

	err := s.TrojanAddUser(ctx, c.Password[0], req.Opt)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	rsp.Index = c.hash()
	rsp.Config = &svcpb.Config{
		Name:       Name,
		ConfigType: svcpb.ConfigType_JSON,
		Config:     c.ExportJson(),
	}
	switch req.ConfigType {
	case svcpb.ConfigType_JSON:
		rsp.Config.ConfigType = svcpb.ConfigType_JSON
		rsp.Config.Config = c.ExportJson()
	case svcpb.ConfigType_URL:
		rsp.Config.ConfigType = svcpb.ConfigType_URL
		rsp.Config.Config = []byte(c.ExportURL())
	default:
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = "unknown config type"
	}
	return
}

func (s *Service) CreateByConfig(ctx context.Context, req *svcpb.CreateByConfigReq) (rsp *svcpb.CreateByConfigRsp, e error) {
	rsp = &svcpb.CreateByConfigRsp{Code: svcpb.Code_OK}

	return
}

func (s *Service) Get(ctx context.Context, req *svcpb.GetReq) (rsp *svcpb.GetRsp, e error) {
	rsp = &svcpb.GetRsp{Code: svcpb.Code_OK}
	userStatus, err := s.TrojanGetUser(ctx, req.Index)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	c := s.configFromStatus(userStatus)
	rsp.Index = c.hash()
	rsp.Config = &svcpb.Config{
		Name:       Name,
		ConfigType: svcpb.ConfigType_JSON,
		Config:     c.ExportJson(),
	}
	return
}

func (s *Service) Delete(ctx context.Context, req *svcpb.DeleteReq) (rsp *svcpb.DeleteRsp, e error) {
	rsp = &svcpb.DeleteRsp{Code: svcpb.Code_OK}
	err := s.TrojanDelUser(ctx, req.Index)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	return
}

func (s *Service) GetAll(ctx context.Context, req *svcpb.GetAllReq) (rsp *svcpb.GetAllRsp, e error) {
	rsp = &svcpb.GetAllRsp{Code: svcpb.Code_OK}
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	stream, err := s.srv.ListUsers(newCtx, &trojanpb.ListUsersRequest{})
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	defer stream.CloseSend()
	rsp.All = make(map[string]*svcpb.Config)
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		c := s.configFromStatus(resp.Status)
		rsp.All[c.hash()] = &svcpb.Config{
			Name:       Name,
			ConfigType: svcpb.ConfigType_JSON,
			Config:     c.ExportJson(),
		}
	}
	return
}

func (s *Service) UpdateOption(ctx context.Context, req *svcpb.UpdateOptionReq) (rsp *svcpb.UpdateOptionRsp, e error) {
	rsp = &svcpb.UpdateOptionRsp{Code: svcpb.Code_OK}
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	setStream, err := s.srv.SetUsers(newCtx)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	defer setStream.CloseSend()
	setReq := &trojanpb.SetUsersRequest{
		Operation: trojanpb.SetUsersRequest_Modify,
		Status: &trojanpb.UserStatus{
			User: &trojanpb.User{
				Hash: req.Index,
			},
		},
	}
	setReq.Status.SpeedLimit = &trojanpb.Speed{
		UploadSpeed:   req.Opt.SendRateLimit,
		DownloadSpeed: req.Opt.RecvRateLimit,
	}
	err = setStream.Send(setReq)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	resp, err := setStream.Recv()
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	if !resp.Success {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = resp.Info
		return
	}
	return
}

func (s *Service) GetStat(ctx context.Context, req *svcpb.GetStatReq) (rsp *svcpb.GetStatRsp, e error) {
	rsp = &svcpb.GetStatRsp{Code: svcpb.Code_OK}
	userStatus, err := s.TrojanGetUser(ctx, req.Index)
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	rsp.Stat = &svcpb.Stat{
		SendTraffic: userStatus.TrafficTotal.UploadTraffic,
		RecvTraffic: userStatus.TrafficTotal.DownloadTraffic,
	}
	return
}

func (s *Service) TrojanGetUser(ctx context.Context, hash string) (*trojanpb.UserStatus, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	getStream, err := s.srv.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	defer getStream.CloseSend()
	err = getStream.Send(&trojanpb.GetUsersRequest{User: &trojanpb.User{Hash: hash}})
	if err != nil {
		return nil, err
	}
	getResp, err := getStream.Recv()
	if err != nil {
		return nil, err
	}
	if getResp.Success != true {
		return nil, errors.New(getResp.Info)
	}
	return getResp.Status, nil
}

func (s *Service) TrojanAddUser(ctx context.Context, password string, opt *svcpb.Option) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	setStream, err := s.srv.SetUsers(ctx)
	if err != nil {
		return err
	}
	defer setStream.CloseSend()
	req := &trojanpb.SetUsersRequest{
		Operation: trojanpb.SetUsersRequest_Add,
		Status: &trojanpb.UserStatus{
			User: &trojanpb.User{
				Password: password,
				Hash:     common.SHA224String(password),
			},
		},
	}
	if opt != nil {
		req.Status.SpeedLimit = &trojanpb.Speed{
			UploadSpeed:   opt.SendRateLimit,
			DownloadSpeed: opt.RecvRateLimit,
		}
	}

	err = setStream.Send(req)
	if err != nil {
		return err
	}
	resp, err := setStream.Recv()
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Info)
	}
	return nil
}

func (s *Service) TrojanDelUser(ctx context.Context, hash string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	setStream, err := s.srv.SetUsers(ctx)
	if err != nil {
		return err
	}
	defer setStream.CloseSend()
	err = setStream.Send(&trojanpb.SetUsersRequest{
		Operation: trojanpb.SetUsersRequest_Delete,
		Status: &trojanpb.UserStatus{
			User: &trojanpb.User{
				Hash: hash,
			},
		},
	})
	if err != nil {
		return err
	}
	resp, err := setStream.Recv()
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Info)
	}
	return nil
}

func (s *Service) configFromStatus(status *trojanpb.UserStatus) *TrojanClientConfig {
	return &TrojanClientConfig{
		RemoteAddr: s.conf.LocalAddr,
		RemotePort: s.conf.LocalPort,
		Password:   []string{status.User.Password},
		SSL: SSL{
			Verify: false,
			SNI:    s.conf.SSL.SNI,
		},
		Mux: Mux{
			Enable:      s.conf.Mux.Enable,
			Concurrency: s.conf.Mux.Concurrency,
			IdleTimeout: s.conf.Mux.IdleTimeout,
		},
		Websocket: Websocket{
			Enable: s.conf.Mux.Enable,
			Path:   s.conf.Websocket.Path,
			Host:   s.conf.Websocket.Host,
		},
	}
}
