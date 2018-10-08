FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/workflow_boot
COPY ./ /go/src/workflow_boot

RUN set -ex && \
go build -v -o /go/bin/workflow_boot \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
COPY --from=builder /go/bin/workflow_boot /usr/bin/
COPY ./bin /workflow/bin
ENTRYPOINT ["/usr/bin/workflow_boot"]
