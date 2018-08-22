FROM node:10.9
MAINTAINER foxzhong <foxzhong@tencent.com>

RUN npm install --global mocha@5.2.0

RUN mkdir -p /root/src
ADD / /root/src

WORKDIR /root/src

CMD ./component-mocha

LABEL TencentHubComponent='{\
  "description": "TencentHub Node.js component, using mocha to run unit test",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "TEST_PATH", "default": "test/*", "desc": "optional, target files path to run mocha test"},\
    {"name": "TEST_PARAMS", "desc": "optional, options pass to mocha cmd, like `--timeout 3000`"}\
  ],\
  "output": []\
}'
