FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-ubuntu-cmd

COPY ./ /go/src/component-ubuntu-cmd

RUN set -ex && \
go build -v -o /go/bin/component-ubuntu-cmd \
-gcflags '-N -l' \
./*.go

FROM ubuntu:16.04
RUN apt-get update && apt-get install -y git
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-ubuntu-cmd /usr/bin/
CMD ["component-ubuntu-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 在Ubuntu环境里clone git代码，并执行用户自定义的shell命令",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码目标提交号或者分支名, 默认为master"},\
    {"name": "CMD", "desc": "必填, 用户自定义shell命令, 支持多行, 使用`/bin/sh -c`执行"}\
  ],\
  "output": [\
  ]\
}'
