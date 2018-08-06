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
  "description": "TencentHub container component, build docker image",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "USERNAME", "desc": "required, kubernetes username"},\
    {"name": "PASSWORD", "desc": "required, kubernetes password"},\
    {"name": "CERTIFICATE", "desc": "required, kubernetes certificate"},\
    {"name": "SERVER", "desc": "required, kubernetes server address"},\
    {"name": "CMD", "desc": "required, the command run in /bin/sh"}\
  ],\
  "output": [\
  ]\
}'
