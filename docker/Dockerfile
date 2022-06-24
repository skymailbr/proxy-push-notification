FROM golang:1.17-bullseye

WORKDIR /talk-push-proxy

RUN apt-get update && apt-get install -y --no-install-recommends \
  ca-certificates \
  libc6 \
  libffi-dev \
  netcat-openbsd \
  tzdata

#RUN cd /talk-push-proxy && \
#  make build-release && \
#  make clean

ENV PUSH_PROXY=/talk-push-proxy/bin/proxy-push-notification

EXPOSE 8066

ENTRYPOINT ["/usr/local/bin/entrypoint"]