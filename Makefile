export GOPROXY=https://goproxy.io,direct
export GOSUMDB=off

BUILD_VERSION   := $(shell git describe --tags)
GIT_COMMIT_SHA1 := $(shell git rev-parse HEAD)
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := simple
VERSION_PACKAGE_NAME := github.com/fregie/PrintVersion

DESCRIBE := Simple server

prebuild:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6.0

go_grpc_opt := --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative
grpc_gw_opt := ${go_grpc_opt} --grpc-gateway_out . --grpc-gateway_opt paths=source_relative
ipaths      := -I. -I./proto/googleapis

protobuf: 
	protoc ${grpc_gw_opt} ${ipaths} proto/api/*.proto
	protoc ${go_grpc_opt} ${ipaths} proto/simple-interface/*.proto

doc:
	protoc ${ipaths} --doc_out=./docs/api --doc_opt=markdown,api.md proto/api/api.proto
	protoc ${ipaths} --doc_out=./docs --doc_opt=html,index.html proto/api/api.proto

adapter-trojan:
	go build -o output/adapter-trojan ./adapter/trojan

adapter-shadowsocks:
	go build -o output/adapter-shadowsocks ./adapter/shadowsocks

simple: protobuf
	go build -ldflags "\
		-X ${VERSION_PACKAGE_NAME}.Version=${BUILD_VERSION} \
		-X '${VERSION_PACKAGE_NAME}.BuildTime=${BUILD_TIME}' \
		-X '${VERSION_PACKAGE_NAME}.GitCommitSHA1=${GIT_COMMIT_SHA1}' \
		-X '${VERSION_PACKAGE_NAME}.Describe=${DESCRIBE}' \
		-X '${VERSION_PACKAGE_NAME}.Name=${BUILD_NAME}'" \
		-o output/simple

.PHONY: spctl
spctl:
	go build -ldflags="-s -w \
		-X ${VERSION_PACKAGE_NAME}.Version=${BUILD_VERSION} \
		-X '${VERSION_PACKAGE_NAME}.BuildTime=${BUILD_TIME}' \
		-X '${VERSION_PACKAGE_NAME}.GitCommitSHA1=${GIT_COMMIT_SHA1}' \
		-X '${VERSION_PACKAGE_NAME}.Describe=${DESCRIBE}' \
		-X '${VERSION_PACKAGE_NAME}.Name=${BUILD_NAME}'" \
		-o output/spctl ./spctl

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

.PHONY: test
test:
	rm -f docker/simple/simple.db
	rm -f docker/trojan-go/data/trojan.db
	docker-compose -f docker/docker-compose-test.yaml up --build -d
	sleep 3
	go test -v ./test -count=1
	docker-compose -f docker/docker-compose-test.yaml down