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

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pb"
)

// NewCmdLeagues returns "trinquetctl leagues" command.
func NewCmdLeagues() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leagues",
		Short: "Print the available leagues",
		RunE: func(cmd *cobra.Command, args []string) error {
			printAvailableLeagues()
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&httpAddress, "httpAddress", "127.0.0.1:8080", "Http address of the gRPC server")

	return cmd
}

func printAvailableLeagues() {
	client, err := getClient(httpAddress)
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	fmt.Println(greenOut("Availables leagues:"))
	resp, err := client.List(context.Background(), &pb.GetLeaguesRequest{})
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	for _, league := range resp.Leagues {
		fmt.Printf("- %s\n", league.Name)
	}

}
