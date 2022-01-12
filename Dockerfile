FROM golang:1.16 AS build-env
MAINTAINER sean bugfan "908958194@qq.com"
ADD . /wireguard-auth
WORKDIR /wireguard-auth
RUN go build -o wireguard-auth main.go

FROM ubuntu:20.04
    
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-env /wireguard-auth/wireguard-auth /wireguard-auth

RUN chmod +x /wireguard-auth

ENTRYPOINT ["/wireguard-auth"]
