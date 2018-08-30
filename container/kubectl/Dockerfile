FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-kubectl

COPY ./ /go/src/component-kubectl

RUN set -ex && \
go build -v -o /go/bin/component-kubectl \
-gcflags '-N -l' \
./*.go

FROM roffe/kubectl
#RUN apk add --update  git
#RUN mkdir -p /root/src
#WORKDIR /root/src
COPY --from=builder /go/bin/component-kubectl /usr/bin/
CMD ["component-kubectl"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 执行 kubectl 命令",\
  "input": [\
    {"name": "USERNAME", "desc": "必填，kubernetes 用户名"},\
    {"name": "PASSWORD", "desc": "必填，kubernetes 用户密码"},\
    {"name": "CERTIFICATE", "desc": "必填, kubernetes 证书内容"},\
    {"name": "SERVER", "desc": "必填, kubernetes 服务器地址"},\
    {"name": "COMMAND", "desc": "必填, CMD命令"}\
  ],\
  "output": [\
  ]\
}'
