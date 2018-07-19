FROM php:latest
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN mkdir -p /root/src
ADD / /root/src

WORKDIR /root/src

CMD ./component-composer
# CMD ./bin/containerops-php Composer

LABEL TencentHubComponent='{\
  "description": "TencentHub PHP component, using Composer to install and upload dependence.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url"},\
    {"name": "GIT_REF", "desc": "optional, target git commit id, git tag or branch name, default is master"},\

    {"name": "HUB_REPO", "desc": "optional,the artifact repo"},\
    {"name": "HUB_USER", "desc": "required, the docker registry user name"},\
    {"name": "HUB_TOKEN", "desc": "required, the docker registry user token"},\
    {"name": "ARTIFACT_TAG", "desc": "optional, the tag of artifact in repo"},\
    {"name": "ARTIFACT_PATH", "desc": "optional, the path of artifact in repo"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT_URL", "desc": "the uploaded artifact url"}\
  ]\
}'

