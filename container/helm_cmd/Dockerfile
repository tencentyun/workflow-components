FROM golang:1.10-alpine as builder

MAINTAINER rockerchen@tencent.com

RUN apk update && apk upgrade && apk add bash && apk add curl && apk add openssl

RUN curl https://raw.githubusercontent.com/helm/helm/master/scripts/get | bash

WORKDIR /go/src/component-helm-cmd

COPY ./ /go/src/component-helm-cmd

RUN set -ex && \
  go build -v -o /go/bin/component-helm-cmd \
  -gcflags '-N -l' \
  ./*.go

FROM roffe/kubectl

RUN apk update && apk upgrade && apk add ca-certificates

COPY --from=builder /go/bin/component-helm-cmd /usr/bin/

COPY --from=builder /usr/local/bin/helm /usr/local/bin/helm


RUN helm init --client-only


CMD ["component-helm-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在预装 helm, kubectl 的环境里执行用户自定义的shell命令",\
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
