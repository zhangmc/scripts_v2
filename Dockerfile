FROM golang:1.17.5-buster

ENV CODE_DIR=/go/src/scripts

COPY /shell/docker-entrypoint.sh /bin/docker-entrypoint

RUN apt update -y \
    && apt install -y bash vim cron git \
    && chsh -s /bin/bash \
    && echo Asia/Shanghai > /etc/timezone && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && export LC_ALL="C.UTF-8" \
    && chmod +x /bin/docker-entrypoint


ENTRYPOINT ["/bin/docker-entrypoint"]

CMD ["/bin/bash"]

WORKDIR $CODE_DIR