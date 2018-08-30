FROM docker.io/phusion/baseimage:0.9.21
MAINTAINER halewang <halewang@tencent.com>

RUN apt-get update && apt-get install -y python3-pip git python-pip python curl

COPY ./bootstrap.py /usr/local/bin/bootstrap.py

WORKDIR /tmp

CMD /usr/local/bin/bootstrap.py