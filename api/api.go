package api

import (
	"context"

	"github.com/fregie/simple/manager"
	svcpb "github.com/fregie/simple/proto/simple-interface"

	proto "github.com/fregie/simple/proto/api"
)

type SimpleAPI struct {
	SessManager *manager.Manager
	proto.UnimplementedSimpleAPIServer
}

func (s *SimpleAPI) CreateSession(ctx context.Context, req *proto.CreateSessionReq) (rsp *proto.CreateSessionRsp, e error) {
	rsp = &proto.CreateSessionRsp{Code: proto.Code_OK}
	sess, err := s.SessManager.CreateSession(ctx, req.Name, req.Proto, req.ConfigType, req.Opt, req.CustomOpt)
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
	sessions := s.SessManager.GetAllSession()
	rsp.Sessions = make([]*proto.Session, 0)
	for _, sess := range sessions {
		s := convertSession(sess)
		rsp.Sessions = append(rsp.Sessions, s)
	}

	return
}

func (s *SimpleAPI) GetSession(ctx context.Context, req *proto.GetSessionReq) (rsp *proto.GetSessionRsp, e error) {
	rsp = &proto.GetSessionRsp{Code: proto.Code_OK}
	var sess *manager.Session
	sess = s.SessManager.GetSession(req.IDorName)
	if sess == nil {
		sess = s.SessManager.GetSessionByName(req.IDorName)
		if sess == nil {
			rsp.Code = proto.Code_InternalError
			rsp.Msg = "session not found"
			return
		}
	}
	rsp.Session = convertSession(sess)
	return
}

func (s *SimpleAPI) DeleteSession(ctx context.Context, req *proto.DeleteSessionReq) (rsp *proto.DeleteSessionRsp, e error) {
	rsp = &proto.DeleteSessionRsp{Code: proto.Code_OK}
	err := s.SessManager.DeleteSession(ctx, req.IDorName)
	if err != nil {
		rsp.Code = proto.Code_InternalError
		rsp.Msg = err.Error()
		return
	}
	return
}

func (s *SimpleAPI) GetProtos(ctx context.Context, req *proto.GetProtosReq) (rsp *proto.GetProtosRsp, e error) {
	rsp = &proto.GetProtosRsp{Code: proto.Code_OK}
	rsp.Protos = s.SessManager.GetProtos()
	return
}

func (s *SimpleAPI) GetSchema(ctx context.Context, req *proto.GetSchemaReq) (rsp *proto.GetSchemaRsp, e error) {
	rsp = &proto.GetSchemaRsp{Code: proto.Code_OK}
	rsp.Schemas = make(map[string]*proto.Schema)

	return
}

func convertSession(s *manager.Session) *proto.Session {
	return &proto.Session{
		ID:         s.ID,
		Name:       s.Name,
		Proto:      s.Proto,
		ConfigType: svcpb.ConfigType(s.ConfigType),
		Config:     s.Config,
		Opt: &svcpb.Option{
			SendRateLimit: s.SendRateLimit,
			RecvRateLimit: s.RecvRateLimit,
		},
	}
}
