FROM golang:alpine AS builder

# Create /duplexity and ~/.ssh
RUN mkdir -p /duplexity ~/.ssh

# Install necessary build tools
RUN apk update && apk upgrade && apk add --no-cache bash git openssh

# Copy repo into build container
COPY . /duplexity

# Move to the client
WORKDIR /duplexity/cmd/websocket
# Build binary in /duplexity called websocket
RUN ["go", "build", "-o", "/duplexity/websocket"]

FROM alpine

# Create /duplexity and move there
WORKDIR /duplexity
# Copy binary from builder
COPY --from=builder /duplexity/websocket .

# Run binary
CMD ./websocket