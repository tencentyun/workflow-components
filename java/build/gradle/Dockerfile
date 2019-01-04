FROM golang:1.10-alpine as builder
MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-gradle
COPY ./ /go/src/component-gradle
RUN set -ex && \
go build -v -o /go/bin/component-gradle \
-gcflags '-N -l' \
./*.go


FROM  gradle:4.9.0-jdk8
USER root
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/component-gradle /usr/bin/
CMD ["component-gradle"]
LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Java Build Gradle), 使用gradle进行java项目构建",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码目标提交号或者分支名, 默认为master"},\
    {"name": "GRADLE_TASK", "desc": "非必填，任务名, 默认任务是 jar"},\
    {"name": "ENTRY_FILE", "desc": "非必填，构建文件, 默认是build.gradle"},\
    {"name": "HUB_REPO", "desc": "非必填，构建产物目标二进制仓库, 组件在构建完成后将尝试自动上传构建产物到此仓库, 如果此值为空将不会自动上传"},\
    {"name": "HUB_USER", "desc": "非必填，上传构建产物对应仓库的用户名, 如果想使用当前主账号, 可以直接设置_WORKFLOW_FLAG_HUB_TOKEN: true, 执行引擎将自动注入当前用户名和token"},\
    {"name": "HUB_TOKEN", "desc": "非必填，上传构建产物对应仓库的用户密码或者token, 同上, 如果设置了_WORKFLOW_FLAG_HUB_TOKEN: true, 此入参可以省略"},\
    {"name": "_WORKFLOW_FLAG_HUB_TOKEN", "default": "true", "desc": "非必填, 若为真, 工作流将根据用户名和密码自动填充HUB_USER和HUB_TOKEN"},\
    {"name": "ARTIFACT_TAG", "desc": "非必填，上传构建产物对应的tag, 默认是latest"},\
    {"name": "ARTIFACT_PATH", "desc": "非必填，上传构建产物对应的仓库目录, 默认是仓库根目录"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT", "desc": "构建产物结果列表"},\
    {"name": "ARTIFACT_URL", "desc": "成功上传的构建产物url"}\
  ]\
}'
