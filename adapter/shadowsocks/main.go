package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	svcpb "github.com/fregie/simple/proto/simple-interface"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcAddr    = flag.String("a", "127.0.0.1:10003", "listen addr")
	ssPortRange = flag.String("p", "50000-50100", "port range")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	minPort, maxPort, err := parsePortRange(*ssPortRange)
	if err != nil {
		log.Fatalf("failed to parse port range: %v", err)
	}
	grpcServer := grpc.NewServer()
	serivce, err := NewService(minPort, maxPort)
	if err != nil {
		log.Fatal(err)
	}
	svcpb.RegisterInterfaceServer(grpcServer, serivce)
	reflection.Register(grpcServer)
	log.Printf("Listening grpc on %s", lis.Addr().String())
	log.Fatal(grpcServer.Serve(lis))
}

// parsePortRange parses a port range string into a min and max port.
func parsePortRange(portRange string) (minPort, maxPort int, err error) {
	_, err = fmt.Sscanf(portRange, "%d-%d", &minPort, &maxPort)
	return
}
