FROM amd64/golang:1.11.2-alpine3.8

MAINTAINER halewang@tencent.com

WORKDIR /go/src

RUN apk add --update git

# Disable cgo so that binaries we build will be fully static.
ENV CGO_ENABLED=0

# Install dep
RUN go get github.com/golang/dep/cmd/dep

# Install glide
RUN go get github.com/Masterminds/glide
ENV GLIDE_HOME /home/user/.glide

COPY ./ /go/src

RUN set -ex && \
go build -v -o /usr/bin/component-gobuild \
-gcflags '-N -l' \
./*.go

RUN rm -rf /go/src/*

CMD ["component-gobuild"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Golang Build), 用以对Golang编写的程序进行构建.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"},\
    {"name": "BUILD_PACKAGE_NAME", "desc": "必填，golang构建项目包的包名，比如github.com/golang/dep"},\
    {"name": "PACKAGE_TARGET", "default":".", "desc": "非必填，main包所的目录, 相对路径, 比如./cmd/dep，默认为."},\
    {"name": "BUILD_VENDOR_CMD", "desc": "非必填, 下载必须的依赖包的vendor目录, 比如go get -v $(go list ./... | grep -v vendor)"},\
    {"name": "OUTPUT", "desc": "非必填, 参数，构建产物的名字"}\
  ],\
  "output": []\
}'