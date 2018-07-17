FROM node:latest
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install -g babel-cli
RUN mkdir -p /root/src
# 不允许全局
RUN cd /root && npm install babel-preset-env

ADD / /root/src
COPY babelrc /root/.babelrc

WORKDIR /root/src
CMD ./component-babel

LABEL TencentHubComponent='{\
  "description": "TencentHub Node.js component, using babel to run code build.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "BUILD_PATH", "desc": "required, target files path to build"},\
    {"name": "BUILD_PARAMS", "desc": "optional, options pass to babel cmd, like `--out-dir public`"}\
  ],\
  "output": []\
}'
