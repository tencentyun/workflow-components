FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-kubectl-apply

COPY ./ /go/src/component-kubectl-apply

RUN set -ex && \
go build -v -o /go/bin/component-kubectl-apply \
-gcflags '-N -l' \
./*.go

FROM roffe/kubectl
#RUN apk add --update  git
#RUN mkdir -p /root/src
#WORKDIR /root/src
COPY --from=builder /go/bin/component-kubectl-apply /usr/bin/
CMD ["component-kubectl-apply"]

LABEL TencentHubComponent='{\
  "description": "TencentHub container component, build docker image",\
  "input": [\
    {"name": "USERNAME", "desc": "required, kubernetes username"},\
    {"name": "PASSWORD", "desc": "required, kubernetes password"},\
    {"name": "CERTIFICATE", "desc": "required, kubernetes certificate"},\
    {"name": "SERVER", "desc": "required, kubernetes server address"},\
    {"name": "RESOURCE", "desc": "required, the definition resource yaml"}\
  ],\
  "output": [\
  ]\
}'
