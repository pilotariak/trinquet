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

	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/pb"
)

var (
	leagueName string
)

// NewCmdLeague returns "trinquetctl league" command.
func NewCmdLeague() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "league",
		Short: "Print information about a league",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newgRPCClient(cmd)
			if err != nil {
				return err
			}
			printLeagueDescription(client)
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&leagueName, "name", "", "Name of the league")
	return cmd
}

func printLeagueDescription(gRPCClient *gRPCClient) {
	if len(leagueName) == 0 {
		fmt.Println(redOut("League name can't be empty"))
		return
	}

	conn, err := gRPCClient.getConn()
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	defer conn.Close()

	client := pb.NewLeagueServiceClient(conn)
	resp, err := client.Get(gRPCClient.getContext(), &pb.GetLeagueRequest{
		Name: leagueName,
	})
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	fmt.Printf(greenOut("Levels:\n"))
	for _, level := range resp.League.Levels {
		fmt.Printf("- %s\n", level.Title)
	}
	fmt.Printf(greenOut("Disciplines:\n"))
	for _, discipline := range resp.League.Disciplines {
		fmt.Printf("- %s\n", discipline.Title)
	}
}
