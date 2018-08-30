FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-update-tke-service

COPY ./ /go/src/component-update-tke-service

RUN set -ex && \
go build -v -o /go/bin/component-update-tke-service \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates
#RUN mkdir -p /root/src
#WORKDIR /root/src
COPY --from=builder /go/bin/component-update-tke-service /usr/bin/
CMD ["component-update-tke-service"]

LABEL TencentHubComponent='{\
  "description": "TencentHub 工作流组件, 更新TKE服务",\
  "input": [\
    {"name": "TENCENTCLOUD_SECRET_ID", "desc": "必填，在云API密钥管理上申请的标识身份的SecretId，一个SecretId对应唯一的SecretKey"},\
    {"name": "TENCENTCLOUD_SECRET_KEY", "desc": "必填，SecretId 对应的唯一SecretKey"},\
    {"name": "REGION", "desc": "必填, 区域参数，用来标识希望操作哪个区域的实例"},\
    {"name": "CLUSTER_ID", "desc": "必填, 服务所在的TKE 集群ID"},\
    {"name": "SERVICE_NAME", "desc": "必填, TKE 服务名"},\
    {"name": "CONTAINERS", "desc": "可选, 新镜像，如果服务中一个实例下有多个容器需要传入此参数，需要一个合法的json字符串, 格式例如 `{\"containerName1\": \"image1\", \"containerName2\": \"image2\"}`"},\
    {"name": "IMAGE", "desc": "可选, 新镜像，如果服务中一个实例下只有一个容器可以传此参数(image和containers二者必填一个)"},\
    {"name": "NAMESPACE", "desc": "可选, kubernetes 服务命名空间, 默认为default"}\
  ],\
  "output": [\
  ]\
}'
