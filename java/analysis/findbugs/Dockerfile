FROM golang:1.10-alpine as builder

MAINTAINER halewang@tencent.com
WORKDIR /go/src/component-findbugs

COPY ./ /go/src/component-findbugs

RUN set -ex && \
go build -v -o /go/bin/component-findbugs \
-gcflags '-N -l' \
./*.go


FROM  openjdk:8-jdk-alpine3.8

RUN apk add --no-cache unzip tar wget git

WORKDIR /opt/gradle
RUN wget https://services.gradle.org/distributions/gradle-3.5-bin.zip \
    && unzip gradle-3.5-bin.zip \
    && rm -rf gradle-3.5-bin.zip

ENV gradle3 /opt/gradle/gradle-3.5/bin
ENV PATH /opt/gradle/gradle-3.5/bin:$PATH

WORKDIR /root
COPY --from=builder /go/src/component-findbugs/findbugs.conf /root/findbugs.conf
COPY --from=builder /go/bin/component-findbugs /usr/bin/

CMD ["component-findbugs"]
LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Java Findbugs Analysis), 用以对Java编写的程序进行分析",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填, 源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"}\
  ],\
  "output": []\
}'