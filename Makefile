export GOPROXY=https://goproxy.io,direct
export GOSUMDB=off

BUILD_VERSION   := $(shell git describe --tags)
GIT_COMMIT_SHA1 := $(shell git rev-parse HEAD)
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := simple
VERSION_PACKAGE_NAME := github.com/fregie/PrintVersion

DESCRIBE := Simple server

protobuf: 
	buf beta mod update
	buf generate

doc:
	buf generate

adapter-trojan:
	go build -o output/adapter-trojan ./adapter/trojan

simple:
	go build -o output/simple