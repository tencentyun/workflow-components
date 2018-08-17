FROM golang:1.10-alpine as builder

MAINTAINER halewang@tencent.com
WORKDIR /go/src/component-golint

COPY ./ /go/src/component-golint

RUN set -ex && \
go build -v -o /go/bin/component-golint \
-gcflags '-N -l' \
./*.go


FROM  golang:1.6
RUN mkdir -p /go/src
WORKDIR /go/src

COPY --from=builder /go/bin/component-golint /usr/bin/
COPY --from=builder /go/src/component-golint/golint /usr/bin
CMD ["component-golint"]

LABEL TencentHubComponent='{\
  "description": "TencentHub Golang component, using golint to run code build.",\
  "input": [\
    {"name": "GIT_CLONE_URL", "desc": "required, git clone url, you can get the authorized git repo url with assignment `${_WORKFLOW_GIT_CLONE_URL}`, which is a global environment variable"},\
    {"name": "GIT_REF", "desc": "optional, git target reference, it can be a git commit, git tag or git branch"},\
    {"name": "LINT_PACKAGE", "desc": "optional, golint target package"},\
    {"name": "LINT_PARAMS", "desc": "optional, params to point output format"}\
  ],\
  "output": []\
}'