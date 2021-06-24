FROM golang:1.16.5-alpine3.13 AS build-env
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/trojan
COPY . .
RUN go mod init github.com/fregie/simple/adapter/trojan && \
  go mod tidy && \
  go build -o adapter-trojan

FROM alpine:latest
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/src/trojan/adapter-trojan /adapter-trojan

ENV ADDR "0.0.0.0:10001"
ENV TROJAN_ADDR "127.0.0.1:1552"

ENTRYPOINT [ "/bin/sh", "-c", "/adapter-trojan -a ${ADDR} -t ${TROJAN_ADDR} -config /server.json" ]