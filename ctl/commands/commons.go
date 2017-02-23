// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/pilotariak/trinquet/auth"
)

var (
	greenOut  = color.New(color.FgGreen).SprintFunc()
	yellowOut = color.New(color.FgYellow).SprintFunc()
	redOut    = color.New(color.FgRed).SprintFunc()

	httpAddress string

	tracerName     string
	zipkinAddress  string
	zipkinPort     int
	appdashAddress string
	appdashPort    int
)

type gRPCClient struct {
	ServerAddress string
	Username      string
	Password      string
}

func newgRPCClient(cmd *cobra.Command) (*gRPCClient, error) {
	serverAddr := cmd.Flag("serverAddr")
	if serverAddr == nil {
		return nil, fmt.Errorf("Parameter 'serverAddr' must not be empty")
	}
	username := cmd.Flag("username")
	if username == nil {
		return nil, fmt.Errorf("Parameter 'username' must not be empty")
	}
	password := cmd.Flag("password")
	if password == nil {
		return nil, fmt.Errorf("Parameter 'password' must not be empty")
	}
	return &gRPCClient{
		ServerAddress: serverAddr.Value.String(),
		Username:      username.Value.String(),
		Password:      password.Value.String(),
	}, nil
}
func (client *gRPCClient) getConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		client.ServerAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
}

func (client *gRPCClient) getContext() context.Context {
	authent := auth.MakeBasicAuth(client.Username, client.Password)
	md := metadata.Pairs("authorization", fmt.Sprintf("basic %s", authent))
	return metadata.NewContext(context.Background(), md)
}
