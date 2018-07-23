FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-docker

COPY ./ /go/src/component-docker

RUN set -ex && \
go build -v -o /go/bin/component-docker \
-gcflags '-N -l' \
./*.go

FROM ubuntu
RUN apt-get -yqq update && apt-get -yqq install docker.io && apt-get -yqq install git
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-docker /usr/bin/
CMD ["component-docker"]

LABEL TencentHubComponent='{\
  "description": "TencentHub container component, build docker image",\
  "input": [\
    {"name": "HUB_USER", "desc": "required, the docker registry user name"},\
    {"name": "HUB_TOKEN", "desc": "required, the docker registry user token"},\
    {"name": "IMAGE", "desc": "optional, the result docker image tag, like `hub.cloud.tencent.com/fox/nodejs`, it can with or without tag"},\
    {"name": "TO_HUB_USER", "desc": "required"},\
    {"name": "TO_HUB_TOKEN", "desc": "required"},\
    {"name": "TO_IMAGE", "desc": "required"}\
  ],\
  "output": [\
    {"name": "IMAGE_ID", "desc": "ID of pushed image tag"},\
    {"name": "IMAGE_DIGEST", "desc": "Image digest of pushed image tag"}\
  ]\
}'

