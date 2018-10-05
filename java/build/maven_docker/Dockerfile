FROM golang:1.10-alpine as builder

MAINTAINER julywind@126.com
WORKDIR /go/src/component-maven

COPY ./ /go/src/component-maven

RUN set -ex && \
go build -v -o /go/bin/component-maven \
-gcflags '-N -l' \
./*.go


FROM  maven:3.5-jdk-8
RUN mkdir -p /root/src
WORKDIR /root/src

RUN apt-get update && apt-get install -y docker \
	&& rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/bin/component-maven /usr/bin/
CMD ["component-maven"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Java Build Maven Docker), 使用maven进行java项目构建,并且可执行docker相关的命令.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码目标提交号或者分支名, 默认为master"},\
    {"name": "M2_SETTINGS_XML", "desc": "非必填，$user/.m2/setting.xml配置文件内容，默认使用maven的全局配置"},\
    {"name": "GOALS", "desc": "非必填，maven 构建目标, 默认是package"},\
    {"name": "POM_PATH", "desc": "非必填，pom 文件相对路径, 默认`./pom.xml`"},\
    {"name": "EXT_COMMAND", "desc": "非必填，GOALS之外的命令, 默认不执行"},\
    {"name": "_WORKFLOW_FLAG_HUB_TOKEN", "default": "true", "desc": "非必填, 若为真, 工作流将根据用户名和密码自动填充HUB_USER和HUB_TOKEN"}\
  ]\
}'
