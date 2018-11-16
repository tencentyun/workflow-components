FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-kubectl-cmd

COPY ./ /go/src/component-kubectl-cmd

RUN set -ex && \
go build -v -o /go/bin/component-kubectl-cmd \
-gcflags '-N -l' \
./*.go

FROM roffe/kubectl
#RUN apk add --update  git
#RUN mkdir -p /root/src
#WORKDIR /root/src
COPY --from=builder /go/bin/component-kubectl-cmd /usr/bin/
CMD ["component-kubectl-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在预装 kubectl 的环境里执行用户自定义的shell命令",\
  "input": [\
    {"name": "USERNAME", "desc": "选填，kubernetes 用户名"},\
    {"name": "PASSWORD", "desc": "选填，kubernetes 用户密码"},\
    {"name": "TOKEN", "desc": "选填，kubernetes 登录token"},\
    {"name": "CERTIFICATE", "desc": "必填, kubernetes 证书内容"},\
    {"name": "SERVER", "desc": "必填, kubernetes 服务器地址"},\
    {"name": "CMD", "desc": "必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行"}\
  ],\
  "output": [\
  ]\
}'
