FROM python:2-alpine 
MAINTAINER halewang <halewang@tencent.com>

RUN apk add --update git && pip install pybuilder==0.11.17
RUN mkdir -p /root/src
ADD / /root/src
WORKDIR /root/src
CMD ./main.py

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 使用 nuitka 构建 Python 项目",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "必填参数, git clone url. 如果工作流已经关联了git项目, 用户可以通过全局环境变量`${_WORKFLOW_GIT_CLONE_URL}`获得clone url"},\
    {"name": "GIT_REF", "desc": "可选参数, git 的目标引用, 可以是git commit、 git tag 或者 git branch, 默认是master"},\
    {"name": "ENTRY_FILE", "desc": "必填，构建入口文件"}\
  ],\
  "output": []\
}'
