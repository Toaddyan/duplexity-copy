#!/usr/bin/env bash

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
