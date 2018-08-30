FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-alpine-cmd

COPY ./ /go/src/component-alpine-cmd

RUN set -ex && \
go build -v -o /go/bin/component-alpine-cmd \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-alpine-cmd /usr/bin/
CMD ["component-alpine-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在Alpine 环境里执行用户自定义的shell命令",\
  "input": [\
    {"name": "CMD", "desc": "必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行"}\
  ],\
  "output": [\
  ]\
}'
