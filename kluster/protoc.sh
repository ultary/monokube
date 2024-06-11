#!/bin/sh

# Protocol Buffer Compiler Installation
#  * https://grpc.io/docs/protoc-installation/
#
# ==== macos ===
# brew update
# brew install protobuf
# brew upgrade protobuf
# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0
# go install github.com/wasilibs/go-protoc-gen-grpc/cmd/protoc-gen-grpc_python@v1.64.2

protoc \
  --proto_path . \
  --go_opt paths=source_relative \
  --go_out ./api/grpc/v1 \
  --go-grpc_opt paths=source_relative \
  --go-grpc_out ./api/grpc/v1 \
  --python_out ../core \
  --pyi_out ../core \
  --grpc_python_out ../core \
  api/grpc/v1/system.proto
