package test_test

import (
	"context"
	"testing"

	pb "github.com/fregie/simple/proto/api"
	pbinf "github.com/fregie/simple/proto/simple-interface"
	"github.com/fregie/simple/test"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type apiSuite struct {
	suite.Suite
	srv pb.SimpleAPIClient
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(apiSuite))
}

func (s *apiSuite) SetupSuite() {
	grpcConn, err := grpc.Dial(test.GrpcAddr, grpc.WithInsecure())
	s.Nil(err)
	s.srv = pb.NewSimpleAPIClient(grpcConn)
	err = test.Reset(context.Background(), s.srv)
	s.Nil(err)
}

func (s *apiSuite) TestGetProtos() {
	rsp, err := s.srv.GetProtos(context.Background(), &pb.GetProtosReq{})
	s.Nil(err)
	s.Equal(2, len(rsp.Protos))
}

func (s *apiSuite) TestAddDelSession() {
	ctx := context.Background()
	rsp1, err := s.srv.GetAllSessions(ctx, &pb.GetAllSessionsReq{})
	s.Nil(err)
	s.Equal(0, len(rsp1.Sessions))
	rsp2, err := s.srv.CreateSession(ctx, &pb.CreateSessionReq{
		Name:       "SS-01",
		Proto:      "ss",
		ConfigType: pbinf.ConfigType_JSON,
	})
	s.Nil(err)
	s.Equalf(pb.Code_OK, rsp2.Code, rsp2.Msg)
	s.NotEmpty(rsp2.Config)
	rsp3, err := s.srv.CreateSession(ctx, &pb.CreateSessionReq{
		Name:       "SS-02",
		Proto:      "trojan",
		ConfigType: pbinf.ConfigType_JSON,
	})
	s.Nil(err)
	s.Equal(pb.Code_OK, rsp3.Code)
	s.NotEmpty(rsp3.Config)
	rsp4, err := s.srv.GetAllSessions(ctx, &pb.GetAllSessionsReq{})
	s.Nil(err)
	s.Equal(2, len(rsp4.Sessions))
	rsp5, err := s.srv.DeleteSession(ctx, &pb.DeleteSessionReq{IDorName: rsp4.Sessions[0].ID})
	s.Nil(err)
	s.Equal(pb.Code_OK, rsp5.Code)
	rsp6, err := s.srv.GetAllSessions(ctx, &pb.GetAllSessionsReq{})
	s.Nil(err)
	s.Equal(1, len(rsp6.Sessions))
}
