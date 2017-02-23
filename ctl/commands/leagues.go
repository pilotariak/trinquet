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

// NewCmdLeagues returns "trinquetctl leagues" command.
func NewCmdLeagues() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leagues",
		Short: "Print the available leagues",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newgRPCClient(cmd)
			if err != nil {
				return err
			}
			printAvailableLeagues(client)
			return nil
		},
	}

	return cmd
}

func printAvailableLeagues(gRPCClient *gRPCClient) {
	conn, err := gRPCClient.getConn()
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	defer conn.Close()

	client := pb.NewLeagueServiceClient(conn)
	fmt.Println(greenOut("Availables leagues:"))
	resp, err := client.List(gRPCClient.getContext(), &pb.GetLeaguesRequest{})
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	for _, league := range resp.Leagues {
		fmt.Printf("- %s\n", league.Name)
	}

}
