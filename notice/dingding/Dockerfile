FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-dingding

COPY ./ /go/src/component-dingding

RUN set -ex && \
go build -v -o /go/bin/component-dingding \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates

COPY --from=builder /go/bin/component-dingding /usr/bin/
CMD ["component-dingding"]

LABEL TencentHubComponent='{\
  "description": "TencentHub Java component, using maven to run code build.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url"},\
    {"name": "GIT_REF", "desc": "optional, target git commit id, git tag or branch name, default is master"},\
    {"name": "GOALS", "desc": "optional, mvn options and goals, default is package"},\
    {"name": "POM_PATH", "desc": "optional, the relative path of pom file"},\

    {"name": "HUB_REPO", "desc": "optional,the artifact repo, the result artifact will be uploaded to the repo if HUB_REPO has value"},\
    {"name": "HUB_USER", "desc": "optional, the tencenthub repo user name. if you want use current user name and token, you can just set `_WORKFLOW_FLAG_HUB_TOKEN: true`, then the workflow engine will inject user name and token"},\
    {"name": "HUB_TOKEN", "desc": "optional, the tencenthub repo user token"},\
    {"name": "ARTIFACT_TAG", "desc": "optional, the tag of artifacts in repo"},\
    {"name": "ARTIFACT_PATH", "desc": "optional, the path of artifacts in repo"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT", "desc": "the output artifact after build"},\
    {"name": "ARTIFACT_URL", "desc": "the uploaded artifact url"}\
  ]\
}'
