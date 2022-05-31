FROM golang:1.18.2-alpine

ENV APP /app

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
      build-base \
      less \
      postgresql-client \
      postgresql-dev \
      tzdata \
      shared-mime-info \
      gcompat
