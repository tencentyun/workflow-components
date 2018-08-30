FROM golang:1.10-alpine as builder
MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-checkstyle
COPY ./ /go/src/component-checkstyle
RUN set -ex && go build -v -o /go/bin/component-checkstyle -gcflags '-N -l' ./*.go

#FROM openjdk:8-jdk-alpine3.7
FROM java:8-alpine
RUN apk update && apk upgrade && apk add git
RUN mkdir -p /root/src
WORKDIR /root/src
COPY ./ /root/src
COPY --from=builder /go/bin/component-checkstyle /usr/bin/
CMD ["component-checkstyle"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用 checkstyle 检查 Java 项目语法规则和代码风格",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填参数, git clone url. 如果工作流已经关联了git项目, 用户可以通过全局环境变量`${_WORKFLOW_GIT_CLONE_URL}`获得clone url"},\
    {"name": "GIT_REF", "desc": "可选参数, git 的目标引用, 可以是git commit、 git tag 或者 git branch, 默认是master"},\
    {"name": "ANALYSIS_OPTIONS", "desc": "可选参数, 传递给 checkstyle 的额外参数, 比如 `--debug`"},\
    {"name": "ANALYSIS_TARGET", "default": ".", "desc": "可选参数, checkstyle 检查的目标文件路径"}\
  ],\
  "output": []\
}'
