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

var (
	leagueName string
)

// NewCmdLeague returns "trinquetctl league" command.
func NewCmdLeague() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "league",
		Short: "Print information about a league",
		RunE: func(cmd *cobra.Command, args []string) error {
			printLeagueDescription()
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&httpAddress, "httpAddress", "127.0.0.1:8080", "Http address of the gRPC server")
	cmd.PersistentFlags().StringVar(&leagueName, "name", "", "Name of the league")
	cmd.PersistentFlags().StringVar(&tracerName, "tracer", "", "OpenTracing tracer to used")
	cmd.PersistentFlags().StringVar(&zipkinAddress, "zipkinAddress", "127.0.0.1", "Zipkin host")
	cmd.PersistentFlags().IntVar(&zipkinPort, "zipkinPort", 9441, "Zipkin port")
	return cmd
}

func printLeagueDescription() {
	if len(leagueName) == 0 {
		fmt.Println(redOut("League name can't be empty"))
		return
	}

	client, err := getClient(httpAddress)
	if err != nil {
		fmt.Println(redOut(err))
		return
	}
	resp, err := client.Get(context.Background(), &pb.GetLeagueRequest{
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
