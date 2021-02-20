FROM golang:alpine AS builder 

RUN mkdir -p /duplexity ~/.ssh
# Copy these files for better caching
COPY go.mod go.sum /duplexity/

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        bash \
        git \
        openssh
# Dynamic Variable for different apps
ARG APPTYPE
ENV APPTYPE ${APPTYPE}
# Copy Entire repo into folder 
COPY . /duplexity
WORKDIR /duplexity/cmd/${APPTYPE}
# Create binary in binary folder 
RUN ["go", "build", "-o", "/duplexity/binary/duplexity"]
RUN ["echo", "$APPTYPE" , ">", "/duplexity/binary/APPTYPE.txt"]
# Create a bare image for binary to live in
FROM alpine

WORKDIR /duplexity
COPY --from=builder /duplexity/binary/ .
ARG APPTYPE
ENV APPTYPE ${APPTYPE}
# By default run binary 
CMD ["/duplexity/duplexity"]
