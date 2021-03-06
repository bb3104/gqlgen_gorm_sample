FROM golang:1.15.0-alpine

RUN apk upgrade --no-cache && \
    apk add --update && \
    apk add --no-cache --virtual build-dependencies build-base && \
    apk add --no-cache --update \
    git \
    bash \
    curl

RUN go version

COPY . /go/src/github.com/bb3104/gqlgen_gorm_sample
WORKDIR /go/src/github.com/bb3104/gqlgen_gorm_sample

ENV GOBIN /go/bin

RUN go get github.com/99designs/gqlgen
