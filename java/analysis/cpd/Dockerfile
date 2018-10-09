FROM golang:1.10-alpine as builder

MAINTAINER halewang@tencent.com
WORKDIR /go/src/component-cpd

COPY ./ /go/src/component-cpd

RUN set -ex && \
go build -v -o /go/bin/component-cpd \
-gcflags '-N -l' \
./*.go


FROM  gradle:4.9.0-jdk8

USER root
WORKDIR /root/src

COPY --from=builder /go/bin/component-cpd /usr/bin/

CMD ["component-cpd"]
LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Java cpd Analysis), 用以生成cpd编码的文件",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "OUT_PUT_XML", "desc": "非必填，分析结果显示，默认为true},\
    {"name": "GIT_REF", "desc": "非必填, 源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"}\
  ],\
  "output": []\
}'