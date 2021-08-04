FROM golang:1.16.6-alpine3.13 AS build-env
RUN apk --no-cache add git
WORKDIR /go/src/shadowsocks
COPY . .
RUN go build -o adapter-shadowsocks ./adapter/shadowsocks

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/src/shadowsocks/adapter-shadowsocks /adapter-shadowsocks

ENV ADDR "0.0.0.0:10003"
ENV PORT_RANGE "50000-50100"

ENTRYPOINT [ "/bin/sh", "-c", "/adapter-shadowsocks -a ${ADDR} -p ${PORT_RANGE}" ]