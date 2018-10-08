FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/workflow_clean_cache
COPY ./ /go/src/workflow_clean_cache

RUN set -ex && \
go build -v -o /go/bin/workflow_clean_cache \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
COPY --from=builder /go/bin/workflow_clean_cache /usr/bin/
ENTRYPOINT ["/usr/bin/workflow_clean_cache"]
