package test_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	ssurl "github.com/fregie/shadowsocks-url"
	pb "github.com/fregie/simple/proto/gen/go/api"
	pbinf "github.com/fregie/simple/proto/gen/go/simple-interface"
	"github.com/fregie/simple/test"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type ssSuite struct {
	suite.Suite
	srv pb.SimpleAPIClient
}

func TestSsSuite(t *testing.T) {
	suite.Run(t, new(ssSuite))
}

func (s *ssSuite) SetupSuite() {
	grpcConn, err := grpc.Dial(test.GrpcAddr, grpc.WithInsecure())
	s.Nil(err)
	s.srv = pb.NewSimpleAPIClient(grpcConn)
	err = test.Reset(context.Background(), s.srv)
	s.Nil(err)
}

func (s *ssSuite) TestAddSession() {
	httpReq, err := http.NewRequest("GET", test.HttpTestUrl, nil)
	s.Nil(err)

	rsp, err := s.srv.CreateSession(context.Background(), &pb.CreateSessionReq{
		Name:       "SS-01",
		Proto:      "ss",
		ConfigType: pbinf.ConfigType_JSON,
		CustomOpt: `{
			"port": 50001,
			"password": "123456",
			"method": "aes-256-gcm"
		}`,
	})
	s.Nil(err)
	s.Equal(pb.Code_OK, rsp.Code)
	s.NotEmpty(rsp.Config)
	conf := &test.SSConfig{}
	err = json.Unmarshal([]byte(rsp.Config), conf)
	s.Nil(err)
	s.Equal(test.SimpleHost, conf.Server)
	s.Equal("123456", conf.Password)
	s.Equal("aes-256-gcm", conf.Method)
	s.Equal(50001, conf.Port)
	rsp2, err := ssurl.SSUrl(fmt.Sprintf("127.0.0.1:%d", conf.Port), conf.Method, conf.Password, httpReq)
	s.Nilf(err, "ssurl.SSUrl: %v", err)
	s.Equal(http.StatusOK, rsp2.StatusCode)
}
