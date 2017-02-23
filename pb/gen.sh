#!/usr/bin/env bash

# This software is confidential and proprietary information of
# Orange Applications for Business. You shall not disclose such Confidential
# Information and shall use it only in accordance with the terms of the
# agreement you entecolors.red into. Unauthorized copying of this file, via any
# medium is strictly prohibited.

# generate the gRPC code
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. *.proto

# generate the JSON interface code
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. health.proto

# generate the swagger definitions
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. health.proto

# generate the JSON interface code
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. league.proto

# generate the swagger definitions
protoc -I/usr/local/include -I. -I$GOPATH/src -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. league.proto

# -----------------------------------------------------------------------

go run swagger/swagger.go . > swagger/api.swagger.json
