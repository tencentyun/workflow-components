FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-alpine-cmd

COPY ./ /go/src/component-alpine-cmd

RUN set -ex && \
go build -v -o /go/bin/component-alpine-cmd \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-alpine-cmd /usr/bin/
CMD ["component-alpine-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub component, run CMD in alpine",\
  "input": [\
    {"name": "CMD", "desc": "required, the command run in /bin/sh"}\
  ],\
  "output": [\
  ]\
}'
