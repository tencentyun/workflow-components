FROM composer:latest
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

    {"name": "HUB_REPO", "desc": "optional,the artifact repo, the result vendor will be uploaded to the repo if HUB_REPO has value"},\
    {"name": "HUB_USER", "desc": "optional, the tencenthub repo user name. if you want use current user name and token, you can just set `_WORKFLOW_FLAG_HUB_TOKEN: true`, then the workflow engine will inject user name and token"},\
    {"name": "HUB_TOKEN", "desc": "optional, the tencenthub repo user token"},\
    {"name": "ARTIFACT_TAG", "desc": "optional, the tag of artifact in repo"},\
    {"name": "ARTIFACT_PATH", "desc": "optional, the path of artifact in repo"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT_URL", "desc": "the uploaded artifact url"}\
  ]\
}'

