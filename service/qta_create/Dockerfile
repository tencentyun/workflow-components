FROM golang:1.10-alpine 

MAINTAINER halewang halewang@tencent.com
WORKDIR /go/src

COPY ./ /go/src

RUN set -ex && \
go build -v -o /usr/bin/component-qta-create \
-gcflags '-N -l' \
./*.go

CMD ["component-qta-create"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(Qta Create), 完成QTA测试实例的创建.",\
  "input": [\
    {"name": "NAME", "desc": "必填，测试计划名称"},\
    {"name": "PRODUCT_PATH", "desc": "必填,  安卓包路径"},\
    {"name": "TEST_REPO_URL", "desc": "必填, 测试代码路径"},\
    {"name": "TESTCASENAME", "desc": "必填, 要执行的测试用例集"}\
  ],\
  "output": [\
    {"name": "_WORKFLOW_TASK_PLAN_ID", "desc": "测试计划ID"}\
  ]\
}'