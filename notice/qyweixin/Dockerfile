FROM golang:1.10-alpine as builder

MAINTAINER halewang@tencent.com
WORKDIR /go/src/component-qyweixin

COPY ./ /go/src/component-qyweixin

RUN set -ex && \
go build -v -o /go/bin/component-qyweixin \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates

COPY --from=builder /go/bin/component-qyweixin /usr/bin/
CMD ["component-qyweixin"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Notice Qyweixin), 使用企业微信发送通知消息.",\
  "input": [\
    {"name": "CORP_ID", "desc": "必填，企业微信上的企业ID"},\
    {"name": "AGENT_ID", "desc": "必填， 企业微信上的创建的应用ID"},\
    {"name": "APP_SECRET", "desc": "必填，企业微信上的应用secret"},\
    {"name": "USERS", "desc": "选填，接受信息的个人用户，多个用户之间通过'|'分割, 注意: USERS、PARTYS、TAGS不能同时为空"},\
    {"name": "PARTYS", "desc": "选填，接受信息的群组用户，多个群组之间通过'|'分割"},\
    {"name": "TAGS", "desc": "选填，接受信息的标签用户，多个标签之间通过'|'分割"},\
    {"name": "MESSAGE", "desc": "选填，发送的信息内容"}\
  ],\
  "output": []\
}'