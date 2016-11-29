// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net"

	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/pilotariak/trinquet/pb"
)

const (
	port = ":8080"
)

func init() {
	grpclog.SetLogger(logrus.StandardLogger())
}

func main() {
	logrus.Infof("[trinquet] Create the gRPC servers")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLeagueServiceServer(s, newLeagueService())
	logrus.Infof("[trinquet] Listen on %s", port)
	s.Serve(lis)
}
