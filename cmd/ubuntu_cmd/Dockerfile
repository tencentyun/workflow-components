FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-ubuntu-cmd

COPY ./ /go/src/component-ubuntu-cmd

RUN set -ex && \
go build -v -o /go/bin/component-ubuntu-cmd \
-gcflags '-N -l' \
./*.go

FROM ubuntu:16.04
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-ubuntu-cmd /usr/bin/
CMD ["component-ubuntu-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub component, run CMD in alpine",\
  "input": [\
    {"name": "CMD", "desc": "required, the command run in /bin/sh"}\
  ],\
  "output": [\
  ]\
}'
