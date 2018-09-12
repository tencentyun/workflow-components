FROM golang:1.10-alpine

MAINTAINER halewang halewang@tencent.com
WORKDIR /go/src/component-qta-run 

COPY ./ /go/src/component-qta-run 

RUN set -ex && \
go build -v -o /usr/bin/component-qta-run \
-gcflags '-N -l' \
./*.go

CMD ["component-qta-run"]

LABEL TencentHubComponent='{\
  "description": "TencentHub官方组件(QTA Run), 运行QTA测试实例，并查询运行结果.",\
  "input": [\
    {"name": "PLAN_ID", "desc": "必填, 测试计划ID"}\
  ],\
  "output": []\
}'