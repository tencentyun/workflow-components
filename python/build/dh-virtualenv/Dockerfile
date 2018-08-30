FROM docker.io/phusion/baseimage:0.9.21
MAINTAINER halewang@tencent.com

RUN apt-get update && apt-get install -y python3-pip git curl python \
            dh-virtualenv devscripts python-virtualenv git equivs

COPY ./bootstrap.py /usr/local/bin/bootstrap.py

WORKDIR /tmp

CMD /usr/local/bin/bootstrap.py