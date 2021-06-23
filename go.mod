module github.com/fregie/simple

go 1.16

require (
	github.com/fregie/PrintVersion v0.1.0
	github.com/fregie/gotool v0.1.7
	github.com/fregie/simple-interface v0.0.0-20210618095204-b0b95059e1a7
	github.com/golang/protobuf v1.5.2
	github.com/p4gefau1t/trojan-go v0.10.4
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/fregie/simple-interface => ./simple-interface
