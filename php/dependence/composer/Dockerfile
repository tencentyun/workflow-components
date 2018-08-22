FROM composer:1.7.2
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN mkdir -p /root/src
ADD / /root/src

WORKDIR /root/src

CMD ./component-composer
# CMD ./bin/containerops-php Composer

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用 Composer 来安装PHP项目依赖, 可选地将依赖上传至二进制仓库.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填参数, git clone url. 如果工作流已经关联了git项目, 用户可以通过全局环境变量`${_WORKFLOW_GIT_CLONE_URL}`获得clone url"},\
    {"name": "GIT_REF", "desc": "可选参数, git 的目标引用, 可以是git commit、 git tag 或者 git branch, 默认是master"},\
    {"name": "HUB_REPO", "desc": "可选参数, 二进制仓库, 如果此参数有值, vendor 结果将上传至此仓库"},\
    {"name": "HUB_USER", "desc": "可选参数, Tencenthub 仓库用户名. 如果希望直接使用当前操作者的身份, 可以直接设置`_WORKFLOW_FLAG_HUB_TOKEN:true`, 工作流将会自动注入HUB_USER 和 HUB_TOKEN"},\
    {"name": "HUB_TOKEN", "desc": "可选参数, Tencenthub 仓库用户名token"},\
    {"name": "ARTIFACT_TAG", "desc": "可选参数, 二进制仓库中, 构建产物的tag, 默认是latest"},\
    {"name": "ARTIFACT_PATH", "desc": "可选参数, 二进制仓库中, 构建产物的路径"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT_URL", "desc": "上传成功的构建产物地址"}\
  ]\
}'

