#!/usr/bin/env bash

# generate the gRPC code
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. service.proto

# generate the JSON interface code
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. service.proto

# generate the swagger definitions
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. service.proto
