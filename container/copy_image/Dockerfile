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
  "description": "TencentHub 工作流组件, 实现镜像从一个仓库到另一个仓库的拷贝",\
  "input": [\
    {"name": "HUB_USER", "desc": "必填，来源镜像仓库用户名"},\
    {"name": "HUB_TOKEN", "desc": "必填，来源镜像仓库用户token"},\
    {"name": "IMAGE_TAG", "desc": "必填, 来源镜像"},\
    {"name": "TO_HUB_USER", "desc": "必填，复制的镜像仓库用户名"},\
    {"name": "TO_HUB_TOKEN", "desc": "必填，复制的镜像仓库用户token"},\
    {"name": "TO_IMAGE", "desc": "必填, 复制的目标镜像"}\
  ],\
  "output": [\
    {"name": "IMAGE_ID", "desc": "复制后新镜像生成的Image ID"},\
    {"name": "IMAGE_DIGEST", "desc": "复制后新镜像生成的Digest"}\
  ]\
}'

