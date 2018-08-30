FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-ubuntu-cmd

COPY ./ /go/src/component-ubuntu-cmd

RUN set -ex && \
go build -v -o /go/bin/component-ubuntu-cmd \
-gcflags '-N -l' \
./*.go

FROM ubuntu:16.04
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-ubuntu-cmd /usr/bin/
CMD ["component-ubuntu-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在Ubuntu 环境里执行用户自定义的shell命令",\
  "input": [\
    {"name": "CMD", "desc": "必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行"}\
  ],\
  "output": [\
  ]\
}'
