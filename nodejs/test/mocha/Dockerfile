FROM node:latest
MAINTAINER foxzhong <foxzhong@tencent.com>

# ENV http_proxy  http://dev-proxy.oa.com:8080
# ENV https_proxy  https://dev-proxy.oa.com:8080
# ENV npm_config_proxy http://dev-proxy.oa.com:8080

RUN npm install --global mocha

RUN mkdir -p /root/src
ADD / /root/src

WORKDIR /root/src

CMD ./component-mocha

LABEL TencentHubComponent='{\
  "description": "TencentHub Node.js component, using mocha to run unit test",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "TEST_PATH", "default": "test/*", "desc": "optional, target files path to run mocha test"},\
    {"name": "TEST_PARAMS", "desc": "optional, options pass to mocha cmd, like `--timeout 3000`"}\
  ],\
  "output": []\
}'
