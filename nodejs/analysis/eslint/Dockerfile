FROM node:10.9
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install -g eslint@5.4.0

RUN mkdir -p /root/src
ADD / /root/src
COPY eslintrc.js /root/.eslintrc.js

WORKDIR /root/src

CMD ./component-eslint

LABEL TencentHubComponent='{\
  "description": "TencentHub Node.js component, using eslint to run code analysis.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "LINT_PATH", "default": ".", "desc": "optional, target files path to run eslint"},\
    {"name": "LINT_PARAMS", "desc": "optional, options pass to eslint cmd, like `--format stylish`"}\
  ],\
  "output": []\
}'
