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

package cmd

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"

	"github.com/pilotariak/trinquet/pb/v1"
	"github.com/pilotariak/trinquet/pkg/cmd/utils"
)

type authCmd struct {
	out io.Writer
}

func newAuthCmd(out io.Writer) *cobra.Command {
	authCmd := &authCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage authentication",
		Long:  "Manage authentication. See subcommands.",
		RunE:  nil,
	}
	authLoginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login into the service",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			if err := authCmd.login(client); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(authLoginCmd)
	return cmd
}

func (cmd authCmd) login(gRPCClient *utils.GRPCClient) error {
	log.Info().Msg("Perform authentication")

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v1.NewAuthServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}
	md := metadata.New(map[string]string{
		"authentication": "true",
	})

	newCtx := metadata.NewIncomingContext(ctx, md)
	md, ok := metadata.FromIncomingContext(newCtx)
	if ok {
		log.Debug().Msgf("Metadata: %s", md)
	}

	resp, err := client.Login(newCtx, &v1.LoginRequest{
		Username: gRPCClient.Username,
		Password: gRPCClient.Password,
	})
	if err != nil {
		return err
	}
	log.Info().Str("auth", name).Msgf("Authentication response: %s", resp)
	if err := ioutil.WriteFile(tokenFile, []byte(resp.Token), 0644); err != nil {
		return err
	}
	fmt.Printf(utils.GreenOut("Authentication succeed\n"))
	return nil
}
