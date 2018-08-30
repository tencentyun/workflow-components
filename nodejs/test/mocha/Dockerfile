FROM node:10.9
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install --global mocha@5.2.0

RUN mkdir -p /root/src
ADD / /root/src

WORKDIR /root/src

CMD ./component-mocha

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用mocha执行单元测试",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填，源代码地址，如为私有仓库需要授权; 如需使用系统关联的git仓库, 可以从系统提供的全局环境变量中获取: ${_WORKFLOW_GIT_CLONE_URL}"},\
    {"name": "GIT_REF", "desc": "非必填，源代码git目标引用，可以是一个git branch, git tag 或者git commit ID, 默认值master"},\
    {"name": "TEST_PATH", "default": "test/*", "desc": "必填，目标文件路径"},\
    {"name": "TEST_PARAMS", "desc": "非必填，运行参数，如 `--timeout 3000`"}\
  ],\
  "output": []\
}'
