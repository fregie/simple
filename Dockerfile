FROM golang:1.16.6-alpine3.13 AS build-env
RUN apk --no-cache add protobuf git make gcc g++ libffi-dev openssl-dev libtool
RUN go install github.com/golang/protobuf/protoc-gen-go@latest && cp /go/bin/protoc-gen-go /usr/bin/
WORKDIR /go/src/github.com/simple
COPY . .

RUN make simple

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/src/github.com/simple/output/simple /simple
COPY --from=build-env /go/src/github.com/simple/example_config.yaml /config.yaml

ENTRYPOINT [ "/simple", "-c", "/config.yaml" ]