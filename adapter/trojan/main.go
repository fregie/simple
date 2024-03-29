package main

import (
	"flag"
	"log"
	"net"

	svcpb "github.com/fregie/simple/proto/simple-interface"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcAddr   = flag.String("a", "127.0.0.1:10001", "listen addr")
	trojanAddr = flag.String("t", "127.0.0.1:2552", "trojan grpc api addr")
	configPath = flag.String("config", "", "trojan-go config path")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	serivce, err := NewService(*configPath, *trojanAddr)
	if err != nil {
		log.Fatal(err)
	}
	svcpb.RegisterInterfaceServer(grpcServer, serivce)
	reflection.Register(grpcServer)
	log.Printf("Listening grpc on %s", lis.Addr().String())
	log.Fatal(grpcServer.Serve(lis))
}
