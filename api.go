package main

import (
	"context"

	svcpb "github.com/fregie/simple/proto/gen/go/simple-interface"

	proto "github.com/fregie/simple/proto/gen/go/api"
)

type SimpleAPI struct {
	proto.UnimplementedSimpleAPIServer
}

func (s *SimpleAPI) CreateSession(ctx context.Context, req *proto.CreateSessionReq) (rsp *proto.CreateSessionRsp, e error) {
	rsp = &proto.CreateSessionRsp{Code: proto.Code_OK}
	sess, err := sessManager.CreateSession(ctx, req.Proto, req.ConfigType, req.Opt, req.CustomOpt)
	if err != nil {
		rsp.Code = proto.Code_InternalError
		rsp.Msg = err.Error()
		return
	}
	rsp.ID = sess.ID
	rsp.Proto = sess.Proto
	rsp.ConfigType = svcpb.ConfigType(sess.ConfigType)
	rsp.Config = sess.Config
	return
}

func (s *SimpleAPI) GetAllSessions(ctx context.Context, req *proto.GetAllSessionsReq) (rsp *proto.GetAllSessionsRsp, e error) {
	rsp = &proto.GetAllSessionsRsp{Code: proto.Code_OK}
	sessions := sessManager.GetAllSession()
	rsp.Sessions = make([]*proto.Session, 0)
	for _, sess := range sessions {
		s := &proto.Session{
			ID:         sess.ID,
			Proto:      sess.Proto,
			ConfigType: svcpb.ConfigType(sess.ConfigType),
			Config:     sess.Config,
			Opt: &svcpb.Option{
				SendRateLimit: sess.SendRateLimit,
				RecvRateLimit: sess.RecvRateLimit,
			},
		}
		rsp.Sessions = append(rsp.Sessions, s)
	}

	return
}

func (s *SimpleAPI) GetSession(ctx context.Context, req *proto.GetSessionReq) (rsp *proto.GetSessionRsp, e error) {
	rsp = &proto.GetSessionRsp{Code: proto.Code_OK}
	sess := sessManager.GetSession(req.ID)
	if sess == nil {
		rsp.Code = proto.Code_InternalError
		rsp.Msg = "session not found"
		return
	}
	rsp.Session = &proto.Session{
		ID:         sess.ID,
		Proto:      sess.Proto,
		ConfigType: svcpb.ConfigType(sess.ConfigType),
		Config:     sess.Config,
		Opt: &svcpb.Option{
			SendRateLimit: sess.SendRateLimit,
			RecvRateLimit: sess.RecvRateLimit,
		},
	}
	return
}

func (s *SimpleAPI) DeleteSession(ctx context.Context, req *proto.DeleteSessionReq) (rsp *proto.DeleteSessionRsp, e error) {
	rsp = &proto.DeleteSessionRsp{Code: proto.Code_OK}
	err := sessManager.DeleteSession(ctx, req.ID)
	if err != nil {
		rsp.Code = proto.Code_InternalError
		rsp.Msg = err.Error()
		return
	}
	return
}

func (s *SimpleAPI) GetProtos(ctx context.Context, req *proto.GetProtosReq) (rsp *proto.GetProtosRsp, e error) {
	rsp = &proto.GetProtosRsp{Code: proto.Code_OK}
	rsp.Protos = sessManager.GetProtos()
	return
}
