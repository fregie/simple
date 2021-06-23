export GOPROXY=https://goproxy.io,direct
export GOSUMDB=off

BUILD_VERSION   := $(shell git describe --tags)
GIT_COMMIT_SHA1 := $(shell git rev-parse HEAD)
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := simple
VERSION_PACKAGE_NAME := github.com/fregie/PrintVersion

DESCRIBE := Simple server

protobuf: 
	protoc --go_out=plugins=grpc,paths=source_relative:. api/*.proto
	protoc --go_out=plugins=grpc,paths=source_relative:. simple-interface/*.proto

adapter-trojan:
	go build -o output/adapter-trojan ./adapter/trojan

simple: protobuf
	go build -o output/simple