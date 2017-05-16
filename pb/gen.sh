#!/usr/bin/env bash

# This software is confidential and proprietary information of
# Orange Applications for Business. You shall not disclose such Confidential
# Information and shall use it only in accordance with the terms of the
# agreement you entecolors.red into. Unauthorized copying of this file, via any
# medium is strictly prohibited.

# generate the gRPC code

function generate_grpcgw {
    pushd $1
    rm -rf *.pb.go
    protoc -I/usr/local/include \
           -I. -I${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --go_out=plugins=grpc:. *.proto

    rm -rf *.pb.gw.go
    protoc -I /usr/local/include -I . \
           -I ${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --grpc-gateway_out=logtostderr=true:. *.proto

    rm -rf ../swagger/*.swagger.json
    protoc -I /usr/local/include -I . \
           -I ${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --swagger_out=logtostderr=true:. *.proto
    popd
}

function generate_grpc {
    pushd $1
    rm -rf *.pb.go
    protoc -I/usr/local/include \
           -I. -I${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --go_out=plugins=grpc:. *.proto
    popd
}


function generate_swagger {
    find . -name "*.json" | xargs -I '{}' mv '{}' swagger/
    rm -f swagger/api.swagger.json
    ls swagger
    go run swagger/swagger.go swagger > swagger/api.swagger.json
}

generate_grpcgw v1beta
generate_grpc health
generate_grpc info

generate_swagger
