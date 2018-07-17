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
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "GIT_TYPE", "desc": "optional, type of `GIT_REF`, it can be branch, tag or commit"},\
    {"name": "IMAGE", "desc": "optional, the result docker image tag, like `hub.cloud.tencent.com/fox/nodejs`, it can with or without tag"},\
    {"name": "IMAGE_TAG_FORMAT", "desc": "optional, image tag format"},\
    {"name": "IMAGE_TAG", "desc": "optional, the result docker image tag, default is `latest`"},\
    {"name": "EXTRA_IMAGE_TAG", "desc": "optional, extra docker image tag"},\
    {"name": "BUILD_WORKDIR", "default": ".", "desc": "optional, the work dir to run docker build"},\
    {"name": "DOCKERFILE_PATH", "default": "Dockerfile", "desc": "optional, Dockerfile path`"},\
    {"name": "BUILD_ARGS", "desc": "optional, arguments pass to docker build as build-arg, it must be valid json string, like `{\"HTTP_PROXY\":\"http://10.20.30.2:1234\",\"TIMEOUT\":\"10\"}`"}\
  ],\
  "output": [\
    {"name": "IMAGE", "desc": "url of pushed image without tag"},\
    {"name": "IMAGE_ID", "desc": "ID of pushed image tag"},\
    {"name": "IMAGE_DIGEST", "desc": "Image digest of pushed image tag"}\
  ]\
}'
