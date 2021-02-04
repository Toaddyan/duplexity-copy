FROM golang:1.15.7-alpine as builder

WORKDIR /duplexity/client
COPY main.go .

RUN ["go","mod","download"]

RUN ["go", "build", "-o", ]