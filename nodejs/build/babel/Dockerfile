FROM node:10.9
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install -g babel-cli@6.26.0
RUN mkdir -p /root/src
# 不允许全局
RUN cd /root && npm install babel-preset-env

ADD / /root/src
COPY babelrc /root/.babelrc

WORKDIR /root/src
CMD ./component-babel

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用 Babel 来构建JavaScript代码, 可选地将构建结果上传至二进制仓库.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填参数, git clone url. 如果工作流已经关联了git项目, 用户可以通过全局环境变量`${_WORKFLOW_GIT_CLONE_URL}`获得clone url"},\
    {"name": "GIT_REF", "desc": "可选参数, git 的目标引用, 可以是git commit、 git tag 或者 git branch, 默认是master"},\
    {"name": "BUILD_PATH", "desc": "必填参数, 需要构建的目标文件路径"},\
    {"name": "OUT_DIR", "desc": "可选参数, babel产物输出目录, 如果为空, 将使用标准输出且不会生成构建产物"},\
    {"name": "BUILD_PARAMS", "desc": "可选参数, 传递给 babel 命令的其他参数, 如`--ignore [list]`"},\
    {"name": "HUB_REPO", "desc": "可选参数, 二进制仓库, 如果此参数有值, 构建结果将上传至此仓库"},\
    {"name": "HUB_USER", "desc": "可选参数, Tencenthub 仓库用户名. 如果希望直接使用当前操作者的身份, 可以直接设置`_WORKFLOW_FLAG_HUB_TOKEN:true`, 工作流将会自动注入HUB_USER 和 HUB_TOKEN"},\
    {"name": "HUB_TOKEN", "desc": "可选参数, Tencenthub 仓库用户名token"},\
    {"name": "ARTIFACT_TAG", "desc": "可选参数, 二进制仓库中, 构建产物的tag, 默认是latest"},\
    {"name": "ARTIFACT_PATH", "desc": "可选参数, 二进制仓库中, 构建产物的路径"}\
  ],\
  "output": [ \
    {"name": "ARTIFACT_URL", "desc": "上传成功的构建产物地址"}\
  ]\
}'
