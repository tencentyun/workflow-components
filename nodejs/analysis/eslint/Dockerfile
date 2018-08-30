FROM node:10.9
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install -g eslint@5.4.0

RUN mkdir -p /root/src
ADD / /root/src
COPY eslintrc.js /root/.eslintrc.js

WORKDIR /root/src

CMD ./component-eslint

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用eslint进行代码分析.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": " 非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"},\
    {"name": "LINT_PATH", "default": ".", "desc": "非必填，目标文件路径"},\
    {"name": "LINT_PARAMS", "desc": " 非必填，运行参数，如 `--format stylish`"}\
  ],\
  "output": []\
}'
