FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-kubectl

COPY ./ /go/src/component-kubectl

RUN set -ex && \
go build -v -o /go/bin/component-kubectl \
-gcflags '-N -l' \
./*.go

FROM roffe/kubectl
#RUN apk add --update  git
#RUN mkdir -p /root/src
#WORKDIR /root/src
COPY --from=builder /go/bin/component-kubectl /usr/bin/
CMD ["component-kubectl"]

LABEL TencentHubComponent='{\
  "description": "TencentHub kubectl component, run CMD with kubectl",\
  "input": [\
    {"name": "USERNAME", "desc": "required, username"},\
    {"name": "PASSWORD", "desc": "required, password"},\
    {"name": "CERTIFICATE", "desc": "required, certificate"},\
    {"name": "SERVER", "desc": "required, server"},\
    {"name": "COMMAND", "desc": "required, command"}\
  ],\
  "output": [\
  ]\
}'
