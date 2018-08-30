FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-kubectl-cmd

COPY ./ /go/src/component-kubectl-cmd

RUN set -ex && \
go build -v -o /go/bin/component-kubectl-cmd \
-gcflags '-N -l' \
./*.go

FROM roffe/kubectl
RUN apk add --update git
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-kubectl-cmd /usr/bin/
CMD ["component-kubectl-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在kubectl环境里clone git代码，并执行用户自定义的shell命令",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码目标提交号或者分支名, 默认为master"},\
    {"name": "USERNAME", "desc": "必填，kubernetes 用户名"},\
    {"name": "PASSWORD", "desc": "必填，kubernetes 用户密码"},\
    {"name": "CERTIFICATE", "desc": "必填, kubernetes 证书内容"},\
    {"name": "SERVER", "desc": "必填, kubernetes 服务器地址"},\
    {"name": "CMD", "desc": "必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行"}\
  ],\
  "output": [\
  ]\
}'
