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

adapter-shadowsocks:
	go build -o output/adapter-shadowsocks ./adapter/shadowsocks

simple:
	go build -o output/simple

.PHONY: spctl
spctl:
	go build -ldflags="-s -w" -o output/spctl ./spctl

.PHONY: docker
docker:
	docker build -t fregie/simple:latest .
	docker build -f ./adapter/trojan/Dockerfile -t fregie/adapter-trojan:latest .
	docker build -f ./adapter/shadowsocks/Dockerfile -t fregie/adapter-shadowsocks:latest .

release:
	tar -czvf output/simple-docker-compose.tar.gz \
		--exclude=docker/simple/simple.db \
		--exclude=docker/trojan-go/data/trojan.db \
		--exclude=docker/docker-compose-build.yaml \
		docker

test:
	docker-compose -f docker/docker-compose-build.yaml up --build -d