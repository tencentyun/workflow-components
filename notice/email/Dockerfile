FROM golang:1.10-alpine as builder

MAINTAINER halewang@tencent.com
WORKDIR /go/src/component-email

COPY ./ /go/src/component-email

RUN set -ex && \
go build -v -o /go/bin/component-email \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates

COPY --from=builder /go/bin/component-email /usr/bin/
COPY --from=builder /go/src/component-email/template.html /usr/bin/
CMD ["component-email"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Notice Email), 使用邮件发送通知消息.",\
  "input": [\
    {"name": "FROM_USER", "desc": "必填，邮件的发送方"},\
    {"name": "TO_USERS", "desc": "必填，邮件的接收方, 如user_xx@qq.com,user_yy@163.com,多个收件人之间通过','分割"},\
    {"name": "SECRET", "desc": "必填，SMTP服务器(邮件发送服务器)的授权码(如QQ和163邮箱)或者邮箱的登录密码(如126邮箱)"},\
    {"name": "SMTP_SERVER_PORT", "desc": "必填，SMTP服务器和端口(smtp.example.com:123),根据发送方的邮箱服务提供商确定SMTP服务器，举例:QQ的SMTP是smtp.qq.com:465、163的SMTP是smtp.163.com:465、126的SMTP是smtp.126.com:25等"},\
    {"name": "SUBJECT", "desc": "选填，邮件的主题"},\
    {"name": "TEXT", "desc": " 选填，邮件的内容"}\
  ],\
  "output": []\
}'