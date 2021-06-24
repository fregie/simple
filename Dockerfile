FROM golang:1.16.5-alpine3.13 AS build-env
RUN apk --no-cache add ca-certificates protobuf git make
RUN go install github.com/golang/protobuf/protoc-gen-go@latest && cp /go/bin/protoc-gen-go /usr/bin/
WORKDIR /go/src/github.com/simple
COPY . .

# RUN git submodule update --init --recursive
RUN make simple

FROM alpine:latest
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/src/github.com/simple/output/simple /simple
COPY --from=build-env /go/src/github.com/simple/example_config.yaml /config.yaml

ENTRYPOINT [ "/simple", "-c", "/config.yaml" ]