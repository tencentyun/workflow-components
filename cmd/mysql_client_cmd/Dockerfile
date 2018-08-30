FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-mysql-client-cmd

COPY ./ /go/src/component-mysql-client-cmd

RUN set -ex && \
go build -v -o /go/bin/component-mysql-client-cmd \
-gcflags '-N -l' \
./*.go

FROM alpine:3.7
RUN apk add --no-cache mysql-client
COPY --from=builder /go/bin/component-mysql-client-cmd /usr/bin/
CMD ["component-mysql-client-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在预装 mysql 客户端的环境里执行用户自定义的shell命令",\
  "input": [\
    {"name": "CMD", "desc": "必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行"}\
  ],\
  "output": [\
  ]\
}'
