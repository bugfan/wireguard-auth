FROM golang:1.16-alpine AS build-env
MAINTAINER sean bugfan "908958194@qq.com"
ADD . /wireguard-auth
WORKDIR /wireguard-auth
RUN go build -o wireguard-auth

FROM alpine:3.13
RUN apk add --update --no-cache
RUN apk  add --update vim && \
    apk  add --update nano
    
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-env /wireguard-auth/wireguard-auth /wireguard-auth

RUN chmod +x /wireguard-auth

ENTRYPOINT ["/wireguard-auth"]
