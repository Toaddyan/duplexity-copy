#!/usr/bin/env bash

# We had some weird trouble trying to get protoc to work. Here's what we did to fix it:
# 827  go get -u github.com/golang/protobuf/protoc-gen-go
# 828  ls
# 829  ./pb/compile-proto.sh 
# 830  clear
# 831  ls
# 832  realpath ~/proto-path/bin/
# 833  vim ~/.bashrc 
# 834  source ~/.bashrc 
# 835  clear
# 836  ls
# 837  pb/compile-proto.sh 
# 838  clear
# 839  ls
# 840  go run main.go 
# 841  go get -u google.golang.org/grpc



# Location of this script
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)" # https://stackoverflow.com/a/246128

# Root of mono-repo
ROOT_DIR=$DIR/../../../

PROTO_DIR=$ROOT_DIR/proto
PB_OUT_DIR=$DIR

# Move to $PROTO_DIR
cd $PROTO_DIR

# Compile gRPC
protoc --go-grpc_out=$PB_OUT_DIR ./backend/v1/backend.proto

# Compile all others
protoc --go_out=$PB_OUT_DIR ./backend/v1/*.proto
