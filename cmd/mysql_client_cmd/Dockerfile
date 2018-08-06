FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-mysql-client-cmd

COPY ./ /go/src/component-mysql-client-cmd

RUN set -ex && \
go build -v -o /go/bin/component-mysql-client-cmd \
-gcflags '-N -l' \
./*.go

FROM alpine:3.7
RUN apk add --no-cache mysql-client
COPY --from=builder /go/bin/component-mysql-client-cmd /usr/bin/
CMD ["component-mysql-client-cmd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub mysql client component, run CMD with mysql client",\
  "input": [\
    {"name": "CMD", "desc": "required, the command run in /bin/sh"}\
  ],\
  "output": [\
  ]\
}'
