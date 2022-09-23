package test

import (
	"context"

	pb "github.com/fregie/simple/proto/api"
)

const (
	GrpcAddr = "127.0.0.1:4433"

	SimpleHost  = "simple.fregie.cn"
	HttpTestUrl = "http://nginx/"
	SqlitePath  = "../docker/simple/simple.db"
)

type SSConfig struct {
	Server   string `json:"server"`
	Port     int    `json:"server_port"`
	Method   string `json:"method"`
	Password string `json:"password"`
}

func Reset(ctx context.Context, srv pb.SimpleAPIClient) error {
	rsp, err := srv.GetAllSessions(ctx, &pb.GetAllSessionsReq{})
	if err != nil {
		return err
	}
	for _, sess := range rsp.Sessions {
		_, err := srv.DeleteSession(ctx, &pb.DeleteSessionReq{IDorName: sess.ID})
		if err != nil {
			return err
		}
	}
	return nil
}
