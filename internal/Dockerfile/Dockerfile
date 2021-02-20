FROM golang:alpine

RUN mkdir -p /duplexity ~/.ssh
COPY go.mod go.sum /duplexity/

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        bash \
        git \
        openssh

WORKDIR /duplexity
RUN ["go","mod", "download"]

