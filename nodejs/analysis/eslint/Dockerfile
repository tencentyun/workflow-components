FROM node:latest
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install eslint -g

RUN mkdir -p /root/src
ADD / /root/src
COPY eslintrc.js /root/.eslintrc.js

WORKDIR /root/src

CMD ./component-eslint

LABEL TencentHubComponent='{\
  "description": "TencentHub Node.js component, using eslint to run code analysis.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "LINT_PATH", "default": ".", "desc": "optional, target files path to run eslint"},\
    {"name": "LINT_PARAMS", "desc": "optional, options pass to eslint cmd, like `--format stylish`"}\
  ],\
  "output": []\
}'
