#!/usr/bin/env bash

# generate the gRPC code

protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. league.proto

# generate the JSON interface code
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. league.proto

# generate the swagger definitions
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:./swagger league.proto

# merge the swagger code into one file
# go run swagger/main.go swagger > ../static/swagger/api.swagger.json
