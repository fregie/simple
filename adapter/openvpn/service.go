package main

import (
	"context"

	pb "github.com/cad/ovpm/api/pb"
	tool "github.com/fregie/gotool"
	svcpb "github.com/fregie/simple/proto/gen/go/simple-interface"
	"google.golang.org/grpc"
)

const Name = "openvpn"

type Service struct {
	srv pb.UserServiceClient
	svcpb.UnimplementedInterfaceServer
}

func NewService(addr string) (*Service, error) {
	m := &Service{}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	m.srv = pb.NewUserServiceClient(conn)
	return m, nil
}

func (s *Service) Name(_ context.Context, _ *svcpb.NameReq) (*svcpb.NameRsp, error) {
	return &svcpb.NameRsp{Name: Name}, nil
}

func (s *Service) IsSupportPersistence(_ context.Context, _ *svcpb.IsSupportPersistenceReq) (*svcpb.IsSupportPersistenceRsp, error) {
	return &svcpb.IsSupportPersistenceRsp{IsSupport: true}, nil
}

func (s *Service) CustomOptionSchema(_ context.Context, _ *svcpb.CustomOptionSchemaReq) (*svcpb.CustomOptionSchemaRsp, error) {
	rsp := &svcpb.CustomOptionSchemaRsp{}
	return rsp, nil
}

func (s *Service) Create(ctx context.Context, req *svcpb.CreateReq) (rsp *svcpb.CreateRsp, e error) {
	rsp = &svcpb.CreateRsp{Code: svcpb.Code_OK}
	c := &OvpnClientConfig{
		Username: string(tool.RandomString(16)),
		Password: string(tool.RandomString(16)),
		NoGw:     false,
		HostID:   0,
	}
	_, err := s.srv.Create(ctx, &pb.UserCreateRequest{
		Username: c.Username,
		Password: c.Password,
		NoGw:     c.NoGw,
		IsAdmin:  false,
	})
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	genRsp, err := s.srv.GenConfig(ctx, &pb.UserGenConfigRequest{Username: c.Username})
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	rsp.Index = c.Index()
	rsp.Config = &svcpb.Config{
		Name:       Name,
		ConfigType: svcpb.ConfigType_TEXT,
		Config:     []byte(genRsp.ClientConfig),
	}

	return
}

func (s *Service) Get(ctx context.Context, req *svcpb.GetReq) (rsp *svcpb.GetRsp, e error) {
	rsp = &svcpb.GetRsp{Code: svcpb.Code_OK}
	genRsp, err := s.srv.GenConfig(ctx, &pb.UserGenConfigRequest{Username: req.Index})
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	rsp.Index = req.Index
	rsp.Config = &svcpb.Config{
		Name:       Name,
		ConfigType: svcpb.ConfigType_TEXT,
		Config:     []byte(genRsp.ClientConfig),
	}
	return
}

func (s *Service) Delete(ctx context.Context, req *svcpb.DeleteReq) (rsp *svcpb.DeleteRsp, e error) {
	rsp = &svcpb.DeleteRsp{Code: svcpb.Code_OK}
	_, err := s.srv.Delete(ctx, &pb.UserDeleteRequest{Username: req.Index})
	if err != nil {
		rsp.Code = svcpb.Code_Fail
		rsp.Msg = err.Error()
		return
	}
	return
}

func (s *Service) GetAll(ctx context.Context, req *svcpb.GetAllReq) (rsp *svcpb.GetAllRsp, e error) {
	rsp = &svcpb.GetAllRsp{Code: svcpb.Code_OK}
	users, err := s.srv.List(ctx, &pb.UserListRequest{})
	if err != nil {
		return nil, err
	}
	rsp.All = make(map[string]*svcpb.Config)
	for _, u := range users.Users {
		rsp.All[u.Username] = &svcpb.Config{
			Name:       Name,
			ConfigType: svcpb.ConfigType_TEXT,
		}
	}
	return
}
