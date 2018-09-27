FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/component-cd

COPY ./ /go/src/component-cd

RUN set -ex && go build -v -o /go/bin/component-cd -gcflags '-N -l' ./*.go

FROM roffe/kubectl
COPY --from=builder /go/bin/component-cd /usr/bin/
CMD ["component-cd"]

LABEL TencentHubComponent='{}'
