FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-update-tke-service

COPY ./ /go/src/component-update-tke-service

RUN set -ex && \
go build -v -o /go/bin/component-update-tke-service \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates
#RUN mkdir -p /root/src
#WORKDIR /root/src
COPY --from=builder /go/bin/component-update-tke-service /usr/bin/
CMD ["component-update-tke-service"]

LABEL TencentHubComponent='{\
  "description": "TencentHub container trigger component, update tke service image",\
  "input": [\
    {"name": "TENCENTCLOUD_SECRET_ID", "desc": "required, tencent cloud secret ID"},\
    {"name": "TENCENTCLOUD_SECRET_KEY", "desc": "required, tencent cloud secret key"},\
    {"name": "REGION", "desc": "required, the region code: Beijing:bj, Guangzhou:gz, Shanghai:sh, HongKong:hk, America:ca, Singapore:sg"},\
    {"name": "CLUSTER_ID", "desc": "required, target cluster ID"},\
    {"name": "SERVICE_NAME", "desc": "required, target service name"},\
    {"name": "CONTAINERS", "desc": "optional, for multiple contanier service, it must be a valid json string, format like `{\"containerName1\": \"image1\", \"containerName2\": \"image2\"}`"},\
    {"name": "IMAGE", "desc": "optonal, for single contanier service, the new image of target service, `CONTAINERS` has high priority"},\
    {"name": "NAMESPACE", "desc": "optional, service namespace, default is `default`"}\
  ],\
  "output": [\
  ]\
}'
