package test_test

import (
	"context"
	"testing"

	"github.com/fregie/simple/manager"
	pb "github.com/fregie/simple/proto/gen/go/api"
	pbinf "github.com/fregie/simple/proto/gen/go/simple-interface"
	"github.com/fregie/simple/test"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	"github.com/cloudquery/sqlite"
	"gorm.io/gorm"
)

type persistenceSuite struct {
	suite.Suite
	srv pb.SimpleAPIClient
}

func TestPersistenceSuite(t *testing.T) {
	suite.Run(t, new(persistenceSuite))
}

func (s *persistenceSuite) SetupSuite() {
	grpcConn, err := grpc.Dial(test.GrpcAddr, grpc.WithInsecure())
	s.Nil(err)
	s.srv = pb.NewSimpleAPIClient(grpcConn)
	err = test.Reset(context.Background(), s.srv)
	s.Nil(err)
}

func (s *persistenceSuite) TestPersistenceSuite() {
	ctx := context.Background()
	db, err := gorm.Open(sqlite.Open(test.SqlitePath), &gorm.Config{})
	s.Nil(err)
	var count int64
	err = db.Model(&manager.Session{}).Count(&count).Error
	s.Nil(err)
	s.Equal(int64(0), count)
	rsp1, err := s.srv.CreateSession(ctx, &pb.CreateSessionReq{
		Name:       "SS-01",
		Proto:      "ss",
		ConfigType: pbinf.ConfigType_JSON,
	})
	s.Nil(err)
	s.Equal(pb.Code_OK, rsp1.Code)
	s.NotEmpty(rsp1.Config)
	sess := &manager.Session{Name: "SS-01"}
	err = db.Where("name = ?", "SS-01").Find(sess).Error
	s.Nil(err)
	s.Equal(rsp1.ID, sess.ID)
	err = db.Model(&manager.Session{}).Count(&count).Error
	s.Nil(err)
	s.Equal(int64(1), count)
	rsp2, err := s.srv.CreateSession(ctx, &pb.CreateSessionReq{
		Name:       "SS-02",
		Proto:      "ss",
		ConfigType: pbinf.ConfigType_JSON,
	})
	s.Nil(err)
	s.Equal(pb.Code_OK, rsp2.Code)
	s.NotEmpty(rsp2.Config)
	sess = &manager.Session{Name: "SS-02"}
	err = db.Where("name = ?", "SS-02").Find(sess).Error
	s.Nil(err)
	s.Equal(rsp2.ID, sess.ID)
	err = db.Model(&manager.Session{}).Count(&count).Error
	s.Nil(err)
	s.Equal(int64(2), count)
	rsp3, err := s.srv.DeleteSession(ctx, &pb.DeleteSessionReq{
		IDorName: rsp1.ID,
	})
	s.Nil(err)
	s.Equal(pb.Code_OK, rsp3.Code)
	err = db.Model(&manager.Session{}).Count(&count).Error
	s.Nil(err)
	s.Equal(int64(1), count)
}
