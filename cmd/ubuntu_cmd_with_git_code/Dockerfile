FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-alpine-cmd

COPY ./ /go/src/component-alpine-cmd

RUN set -ex && \
go build -v -o /go/bin/component-alpine-cmd \
-gcflags '-N -l' \
./*.go

FROM ubuntu:16.04
RUN apk add --update git
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-alpine-cmd /usr/bin/
CMD ["component-alpine-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub component, clone git code and run CMD in alpine",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "CMD", "desc": "required, the command run in /bin/sh"}\
  ],\
  "output": [\
  ]\
}'
