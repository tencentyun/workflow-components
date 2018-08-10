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
  "description": "TencentHub Java component, using maven to run code build.",\
  "input": [\
    {"name": "CORPID", "desc": "required, company ID"},\
    {"name": "APPSECRET", "desc": "optional, application secrect, after admin create an app"},\
    {"name": "AGENTID", "desc": "required, APP ID"},\
    {"name": "USERS", "desc": "optional, the users that admin want to send msg to, note: USERS OR PARTYS OR TAGS you must give one"},\
    {"name": "PARTYS", "desc": "optional, the partys that admin want to send msg to"},\
    {"name": "TAGS", "desc": "optional, the tags that admin want to send msg to"},\
  ],\
  "output": []\
}'