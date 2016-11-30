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
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pilotariak/trinquet/pb"
)

const (
	port = "localhost:8080"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLeagueServiceClient(conn)
	glog.Infoln("[trinquet] Retrieve all leagues")
	resp, err := client.List(context.Background(), &pb.GetLeaguesRequest{})
	if err != nil {
		glog.Fatalf("[trinquet] Could not retrieve leagues: %v", err)
	}
	glog.Infof("[trinquet] Available leagues: %s", resp)
}
