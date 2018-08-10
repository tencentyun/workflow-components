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
    {"name": "AT_MOBILES", "desc": "optional, TODO"},\
    {"name": "IS_AT_ALL", "desc": "optional, TODO"},\
    {"name": "MESSAGE", "desc": "optional, TODO"},\
    {"name": "_WORKFLOW_TASK_DETAIL", "desc": "optional,TODO"},\
    {"name": "_WORKFLOW_FLAG_TASK_DETAIL", "default": "true", "desc": "optional, TODO"}\
  ],\
  "output": [ \
  ]\
}'
