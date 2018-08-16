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
  "description": "TencentHub Java component, using maven to run code build.",\
  "input": [\
    {"name": "FROM_USER", "desc": "requied, the user who post email"},\
    {"name": "TO_USERS", "desc": "requied, the users who accept email, for example, user_xx@qq.com | user_yy@163.com | user_zz@126.com"},\
    {"name": "SECRET", "desc": "requied, the authorization code of SMTP server or login password of mailbox"},\
    {"name": "SMTP_SERVER_PORT", "desc": "required, the SMTP server, such as, smtp.qq.com, smtp.163.com, smtp.126.com"},\
    {"name": "SUBJECT", "desc": "optional, the subject of email"},\
    {"name": "TEXT", "desc": "optional, the content of email"},\
    {"name": "_WORKFLOW_FLAG_TASK_DETAIL", "default": "true", "desc": "optional, TODO"}\
  ],\
  "output": []\
}'