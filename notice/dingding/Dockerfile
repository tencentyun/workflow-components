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
  "description": "TencentHub官方组件(Notice Dingding), 使用钉钉发送通知消息.",\
  "input": [\
    {"name": "WEBHOOK", "desc": "必填, 钉钉机器人Webhook地址"},\
    {"name": "AT_MOBILES", "desc": "非必填，被@人的手机号"},\
    {"name": "IS_AT_ALL", "desc": "非必填，@所有人时:true, 否则为:false"},\
    {"name": "MESSAGE", "desc": "非必填，自定发送的文本消息"}\
  ],\
  "output": [ \
  ]\
}'
