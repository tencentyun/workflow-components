FROM docker.io/phusion/baseimage:0.9.21
MAINTAINER halewang@tencent.com

RUN apt-get update && apt-get install -y python3-pip git python-pip python

RUN pip3 install pyyaml

COPY ./bootstrap.py /usr/local/bin

WORKDIR /tmp

CMD /usr/local/bin/bootstrap.py