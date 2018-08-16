FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-dingding

COPY ./ /go/src/component-dingding

RUN set -ex && \
go build -v -o /go/bin/component-dingding \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates

COPY --from=builder /go/bin/component-dingding /usr/bin/
CMD ["component-dingding"]

LABEL TencentHubComponent='{\
  "description": "TencentHub notice component, using dingding to send message.",\
  "input": [\
    {"name": "WEBHOOK", "desc": "required, dingding webhook url"},\
    {"name": "AT_MOBILES", "desc": "optional, @ user phone in custom notice message"},\
    {"name": "IS_AT_ALL", "desc": "optional, @all in custom notice message"},\
    {"name": "MESSAGE", "desc": "optional, custom notice message"},\
    {"name": "_WORKFLOW_FLAG_TASK_DETAIL", "default": "true", "desc": "optional, auto send flow task detail"}\
  ],\
  "output": [ \
  ]\
}'
