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
CMD ["component-email"]

LABEL TencentHubComponent='{\
  "description": "TencentHub Java component, using maven to run code build.",\
  "input": [\
    {"name": "FROM_USER", "desc": "requied, the user who post email"},\
    {"name": "SECRET", "desc": "requied, the secret"},\
    {"name": "TO_USERS", "desc": "requied, the users who get email"},\
    {"name": "SUBJECT", "desc": "optional, the subject of email"},\
    {"name": "TEXT", "desc": "optional, the content of email"},\
  ],\
  "output": []\
}'