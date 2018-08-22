FROM python:3-alpine
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN apk add --update git && pip install coala-bears==0.10.0
RUN mkdir -p /root/src
ADD / /root/src
WORKDIR /root/src
CMD ./main.py

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用 Coala 检查 Python 项目语法规则和代码风格",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填参数, git clone url. 如果工作流已经关联了git项目, 用户可以通过全局环境变量`${_WORKFLOW_GIT_CLONE_URL}`获得clone url"},\
    {"name": "GIT_REF", "desc": "可选参数, git 的目标引用, 可以是git commit、 git tag 或者 git branch, 默认是master"},\
    {"name": "FILES", "default": "./**/*.py", "desc": "非必填，目标文件, 默认是项目下所有py文件"},\
    {"name": "BEARS", "default": "PEP8Bear,PyUnusedCodeBear", "desc": "非必填，coala bears"}\
  ],\
  "output": []\
}'
