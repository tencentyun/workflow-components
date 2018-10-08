FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-docker

COPY ./ /go/src/component-docker

RUN set -ex && \
go build -v -o /go/bin/component-docker \
-gcflags '-N -l' \
./*.go

FROM ubuntu:18.04
RUN apt-get -yqq update && apt-get -yqq install docker.io && apt-get -yqq install git
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-docker /usr/bin/
CMD ["component-docker"]

LABEL TencentHubComponent='{\
  "description": "TencentHub container component, build docker image",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"},\
    {"name": "GIT_TYPE", "desc": "非必填, 标识`GIT_REF`的类型: `branch`, `tag` 或者 `commit`"},\
    {"name": "IMAGE_TAG_FORMAT", "desc": "非必填, 新增镜像Tag的命名格式, 可以使用git 分支名/tag名/commit ID/当前时间等作为命名组成, 如`latest-$branch-$commit-$time`"},\
    {"name": "IMAGE_TAG", "desc": "非必填, 目标镜像Tag, 默认值`latest`"},\
    {"name": "EXTRA_IMAGE_TAG", "desc": "非必填, 新增目标镜像Tag"},\
    {"name": "BUILD_WORKDIR", "default": ".", "desc": "非必填, 工作路径"},\
    {"name": "DOCKERFILE_PATH", "default": "Dockerfile", "desc": "非必填，Dockerfile路径"},\
    {"name": "NO_CACHE", "default": "false", "desc": "非必填，默认值是false, 如果为true 将开启`--no-cache`禁用 docker 构建缓存标志"},\
    {"name": "BUILD_ARGS", "desc": "非必填，传递给`--build-arg`的构建参数, 必须是一个有效的json字符串"}\
  ],\
  "output": [\
    {"name": "IMAGE", "desc": "构建的镜像地址, 不包含tag名称,"},\
    {"name": "IMAGE_TAG", "desc": "镜像tag名称"},\
    {"name": "IMAGE_WITH_TAG", "desc": "构建的镜像地址, 包含tag名称"},\
    {"name": "IMAGE_ID", "desc": "构建镜像生成的Image ID"},\
    {"name": "IMAGE_DIGEST", "desc": "构建镜像生成的Digest"}\
  ]\
}'
