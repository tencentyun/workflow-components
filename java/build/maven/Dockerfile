FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-maven

COPY ./ /go/src/component-maven

RUN set -ex && \
go build -v -o /go/bin/component-maven \
-gcflags '-N -l' \
./*.go


FROM  maven:3.5-jdk-8
RUN mkdir -p /root/src
WORKDIR /root/src

COPY --from=builder /go/bin/component-maven /usr/bin/
CMD ["component-maven"]

LABEL TencentHubComponent='{\
  "description": "TencentHub Java component, using maven to run code build.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "GOALS", "desc": "optional, mvn options and goals, default is package"},\
    {"name": "POM_PATH", "desc": "optional, the relative path of pom file, default is `./pom.xml`"},\
    {"name": "HUB_REPO", "desc": "optional,the artifact repo, the result artifact will be uploaded to the repo if HUB_REPO has value"},\
    {"name": "HUB_USER", "desc": "optional, the tencenthub repo user name. if you want use current user name and password, you can just set `_WORKFLOW_FLAG_HUB_TOKEN: true`, then the workflow engine will inject user name and token"},\
    {"name": "HUB_TOKEN", "desc": "optional, the tencenthub repo user password"},\
    {"name": "_WORKFLOW_FLAG_HUB_TOKEN", "default": "true", "desc": "optional, if it is true, workflow engine will inject user name and password to `HUB_USER` and `HUB_TOKEN`"},\
    {"name": "ARTIFACT_TAG", "desc": "optional, the tag of artifacts in repo, default is `latest`"},\
    {"name": "ARTIFACT_PATH", "desc": "optional, the path of artifacts in repo"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT", "desc": "the output artifact after build"},\
    {"name": "ARTIFACT_URL", "desc": "the uploaded artifact url"}\
  ]\
}'
