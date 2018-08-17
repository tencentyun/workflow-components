FROM golang:1.10-alpine as builder

MAINTAINER halewang@tencent.com
WORKDIR /go/src

COPY ./ /go/src

RUN set -ex && \
go build -v -o /usr/bin/component-gotest \
-gcflags '-N -l' \
./*.go

RUN apk add --update git

CMD ["component-gotest"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Golang Gotest Analysis), 用以对Golang编写的程序进行测试.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"},\
    {"name": "GTEST_PACKAGE_FILE", "desc": "非必填, 待测试的代码包(多个代码包通过空格分割)或者文件(test文件和源码文件通过空格分割), 代码包通过路径的形式给出, 默认检索所有的代码包"},\
    {"name": "GTEST_PARAMS", "desc": "非必填, 参数，用于指明代码测试信息的输出格式"}\
  ],\
  "output": []\
}'