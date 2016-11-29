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
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/pilotariak/trinquet/pb"
)

const (
	port = "localhost:8080"
)

func init() {
	grpclog.SetLogger(logrus.StandardLogger())
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLeagueServiceClient(conn)
	logrus.Info("[trinquet] Retrieve all leagues")
	resp, err := client.GetLeagues(context.Background(), &pb.GetLeaguesRequest{})
	if err != nil {
		logrus.Fatalf("[trinquet] Could not retrieve leagues: %v", err)
	}
	logrus.Infof("[trinquet] Available leagues: %s", resp)
}
