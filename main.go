package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	svcpb "github.com/fregie/simple/proto/simple-interface"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	version "github.com/fregie/PrintVersion"
	"github.com/fregie/simple/api"
	"github.com/fregie/simple/manager"
	pb "github.com/fregie/simple/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	confFile    = flag.String("c", "/etc/simple.yaml", "config file path")
	showVersion = flag.Bool("version", false, "Displays version and exit.")
	isDebug     = flag.Bool("d", false, "debug mode")
)

var (
	//Debug print debug informantion
	Debug *log.Logger
	//Info print Info informantion
	Info *log.Logger
	//Error print Error informantion
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(ioutil.Discard, "[Debug] ", log.Ldate|log.Ltime|log.Lshortfile)
}

var sessManager *manager.Manager

func main() {
	flag.Parse()
	if *showVersion {
		version.PrintVersion()
		os.Exit(0)
	}
	if *isDebug {
		Debug.SetOutput(os.Stdout)
	}
	conf, err := parseConfigFromFile(*confFile)
	if err != nil {
		Error.Fatalf("Load config failed: %s", err)
	}
	sessManager, err = manager.NewManager(conf.Sqlite, conf.Host)
	if err != nil {
		Error.Fatal(err)
	}

	for _, addr := range conf.Services {
		Info.Printf("Register to service %s", addr)
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			Error.Printf("Dial to service %s failed:%s", addr, err)
			continue
		}
		srv := svcpb.NewInterfaceClient(conn)
		err = sessManager.RegisterService(srv)
		if err != nil {
			Error.Printf("Register service %s failed: %s", addr, err)
			continue
		}
	}
	serverOpts := make([]grpc.ServerOption, 0)
	if conf.PromAddr != "" {
		go runPromHttp(conf.PromAddr)
		serverOpts = append(serverOpts,
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		)
	}
	lis, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		Error.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(serverOpts...)
	pb.RegisterSimpleAPIServer(grpcServer, &api.SimpleAPI{SessManager: sessManager})
	reflection.Register(grpcServer)
	if conf.PromAddr != "" {
		grpc_prometheus.Register(grpcServer)
	}
	if conf.GrpcGatewayAddr != "" {
		_, port, _ := net.SplitHostPort(lis.Addr().String())
		go runGateway(fmt.Sprintf("127.0.0.1:%s", port), conf.GrpcGatewayAddr)
	}
	Info.Printf("Listening grpc on %s", lis.Addr().String())
	grpcServer.Serve(lis)
}

func runGateway(grpcServerEndpoint string, addr string) error {
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers:  true,
			EmitUnpopulated: true,
		},
	}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterSimpleAPIHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	Info.Printf("Listening grpc gateway on %s (proxy to %s)", addr, grpcServerEndpoint)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(addr, mux)
}
