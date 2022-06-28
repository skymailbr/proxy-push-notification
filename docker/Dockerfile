FROM golang:1.17-bullseye

WORKDIR /talk-push-proxy

RUN apt-get update && apt-get install -y --no-install-recommends \
  ca-certificates \
  libc6 \
  libffi-dev \
  netcat-openbsd \
  tzdata

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8066

RUN air init
CMD air