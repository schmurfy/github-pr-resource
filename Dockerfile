FROM alpine:3.15 as resource

COPY build /opt/resource

RUN apk add --update --no-cache \
  curl \
  git \
  git-lfs \
  openssh \
  && chmod +x /opt/resource/*

COPY scripts/askpass.sh /usr/local/bin/askpass.sh
