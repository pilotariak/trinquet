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
	"errors"
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/pb/v1beta"
	"github.com/pilotariak/trinquet/pkg/cmd/utils"
)

var (
	name string
)

type leagueCmd struct {
	out io.Writer
}

func newLeagueCmd(out io.Writer) *cobra.Command {
	leagueCmd := &leagueCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "league",
		Short: "Manage leagues",
		Long:  "Manage leagues. See List, Get, ... subcommands.",
		RunE:  nil,
	}
	listLeagueCmd := &cobra.Command{
		Use:   "list",
		Short: "List all leagues",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			if err := leagueCmd.listLeagues(client); err != nil {
				return err
			}
			return nil
		},
	}
	getLeagueCmd := &cobra.Command{
		Use:   "get",
		Short: "Retreive a league",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(name) == 0 {
				return errors.New("missing league name")
			}
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			if err := leagueCmd.getLeague(client, name); err != nil {
				return err
			}
			return nil
		},
	}

	getLeagueCmd.PersistentFlags().StringVar(&name, "name", "", "League's name")
	cmd.AddCommand(listLeagueCmd)
	cmd.AddCommand(getLeagueCmd)
	return cmd
}

func (cmd leagueCmd) listLeagues(gRPCClient *utils.GRPCClient) error {
	log.Info().Msg("List all leagues")

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v1beta.NewLeagueServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName, tokenFile)
	if err != nil {
		return err
	}

	resp, err := client.List(ctx, &v1beta.GetLeaguesRequest{})
	if err != nil {
		return err
	}
	for _, league := range resp.Leagues {
		fmt.Printf("- %s\n", league.Name)
	}
	return nil
}

func (cmd leagueCmd) getLeague(gRPCClient *utils.GRPCClient, name string) error {
	log.Info().Str("league", name).Msg("Retrieve league informations")

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v1beta.NewLeagueServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}
	resp, err := client.Get(ctx, &v1beta.GetLeagueRequest{
		Name: name,
	})
	if err != nil {
		return err
	}
	fmt.Printf(utils.GreenOut("Levels:\n"))
	for _, level := range resp.League.Levels {
		fmt.Printf("- %s\n", level.Title)
	}
	fmt.Printf(utils.GreenOut("Disciplines:\n"))
	for _, discipline := range resp.League.Disciplines {
		fmt.Printf("- %s\n", discipline.Title)
	}
	return nil
}
