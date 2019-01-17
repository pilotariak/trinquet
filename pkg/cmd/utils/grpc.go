// Copyright (C) 2016-2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	// "bytes"
	"errors"
	"fmt"
	// "net"
	// "os"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	// "github.com/pilotariak/trinquet/pkg/auth"
	_ "github.com/pilotariak/trinquet/pkg/auth/basic"
	_ "github.com/pilotariak/trinquet/pkg/auth/jwt"
	// "github.com/pilotariak/trinquet/pkg/config"
	_ "github.com/pilotariak/trinquet/pkg/credentials/text"
	_ "github.com/pilotariak/trinquet/pkg/credentials/vault"
	"github.com/pilotariak/trinquet/pkg/transport"
)

var (
	ErrUsernameNotFound    = errors.New("Username not found")
	ErrApiKeyNotFound      = errors.New("API key not found")
	ErrGrpcAddressNotFound = errors.New("gRPC address not found")
)

type GRPCClient struct {
	ServerAddress string
	Username      string
	Password      string
	AuthSystem    string
	// Authentication auth.Authentication
}

func NewGRPCClient(cmd *cobra.Command) (*GRPCClient, error) {
	setupFromEnvironmentVariables()
	if len(username) == 0 {
		return nil, ErrUsernameNotFound
	}
	if len(password) == 0 {
		return nil, ErrApiKeyNotFound
	}
	if len(serverAddress) == 0 {
		return nil, ErrGrpcAddressNotFound
	}
	// conf := &config.Configuration{
	// 	Credentials: &config.CredentialsConfiguration{
	// 		Text: &config.TextCredentialsConfiguration{
	// 			Username: username,
	// 			Password: password,
	// 		},
	// 	},
	// 	Auth: &config.AuthConfiguration{
	// 		Name: authSystem,
	// 		// Vault: &config.VaultConfiguration{
	// 		// 	Address: "http://github.com/pilotariak",
	// 		// },
	// 		BasicAuth: &config.BasicAuthConfiguration{},
	// 	},
	// }
	// authentication, err := auth.New(conf)
	// if err != nil {
	// 	return nil, err
	// }
	glog.V(2).Infof("gRPC client created: %s %s", serverAddress, username)
	return &GRPCClient{
		ServerAddress: serverAddress,
		Username:      username,
		Password:      password,
		AuthSystem:    authSystem,
		// Authentication: authentication,
	}, nil
}

func (client *GRPCClient) GetConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		client.ServerAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
}

func (client *GRPCClient) GetContext(cliName string, tokenFile ...string) (context.Context, error) {
	ctx := context.Background()
	// token, err := client.Authentication.Encode(ctx, client.Username, client.Password)
	// if err != nil {
	// 	return nil, err
	// }
	// headers := map[string]string{
	// 	transport.Authorization: auth.GetAuthenticationHeader(client.Authentication, token),
	// }
	// if host, err := os.Hostname(); err != nil {
	// 	headers[transport.UserHostname] = host
	// }
	// addrs, _ := net.InterfaceAddrs()
	// var buffer bytes.Buffer
	// for _, a := range addrs {
	// 	if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			buffer.WriteString(ipnet.IP.String() + " ")
	// 		}
	// 	}
	// }
	// headers[transport.UserIP] = buffer.String()
	// headers[transport.UserID] = client.Username

	headers := map[string]string{}
	if len(tokenFile) > 0 {
		token, err := ioutil.ReadFile(tokenFile[0])
		if err != nil {
			return nil, err
		}
		authorization, err := authenticationHeader(client.AuthSystem, string(token))
		if err != nil {
			return nil, err
		}
		headers[transport.Authorization] = authorization
	}
	md := metadata.New(headers)
	glog.V(2).Infof("Transport metadata: %s", md)
	return metautils.NiceMD(md).ToOutgoing(ctx), nil
}

func authenticationHeader(authSystem string, token string) (string, error) {
	switch authSystem {
	case "JWT":
		return fmt.Sprintf("bearer %s", token), nil
	case "BasicAuth":
		return fmt.Sprintf("basic %s", token), nil
	default:
		return "", fmt.Errorf("Invalid authentication system: %s", authSystem)
	}
}
