FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-maven

COPY ./ /go/src/component-maven

RUN set -ex && \
go build -v -o /go/bin/component-maven \
-gcflags '-N -l' \
./*.go


FROM  maven:3.5-jdk-8
RUN mkdir -p /root/src
WORKDIR /root/src

COPY --from=builder /go/bin/component-maven /usr/bin/
CMD ["component-maven"]

LABEL TencentHubComponent='{\
  "description": "TencentHub Java component, using maven to run code build.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url"},\
    {"name": "GIT_REF", "desc": "optional, target git commit id, git tag or branch name, default is master"},\
    {"name": "GOALS", "desc": "optional, mvn options and goals, default is package"},\
    {"name": "POM_PATH", "desc": "optional, the relative path of pom file"}\
  ],\
  "output": [ \
    {"name": "ARTIFACTS", "desc": "the output artifact after build"}\
  ]\
}'
