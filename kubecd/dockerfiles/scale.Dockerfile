FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/kubecd

COPY ./ /go/src/kubecd

RUN set -ex && go build -v -o /go/bin/kubecd -gcflags '-N -l' ./*.go

FROM roffe/kubectl
COPY --from=builder /go/bin/kubecd /usr/bin/
ENV ACTION scale
CMD ["kubecd"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 基于kubernetes 持续部署组件: 伸缩操作",\
  "input": [\
    {"name": "USERNAME", "desc": "必填，kubernetes 用户名"},\
    {"name": "PASSWORD", "desc": "必填，kubernetes 用户密码"},\
    {"name": "CERTIFICATE", "desc": "必填, kubernetes 证书内容"},\
    {"name": "SERVER", "desc": "必填, kubernetes 服务器地址"},\

    {"name": "DEPLOY_GROUP", "desc": "必填, 目标部署组"},\
    {"name": "DEPLOY_TARGET", "desc": "可选, 目标部署版本游标"},\
    {"name": "DEPLOYMENT_NAME", "desc": "可选, 目标部署版本名称"},\

    {"name": "SCALE_TO", "desc": "可选, 最终副本数"},\
    {"name": "SCALE_UP", "desc": "可选, 向上扩容数量"},\
    {"name": "SCALE_DOWN", "desc": "可选, 向下缩容数量"},\

    {"name": "AUTO_DELETION", "desc": "可选, 当副本数量收缩为0时, 是否删除该版本"}\
  ],\
  "output": [\
  ]\
}'
