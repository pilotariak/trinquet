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

package cmd

import (
	"io"

	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/cmd/utils"
	"github.com/pilotariak/trinquet/pb/health"
)

type healthCmd struct {
	out io.Writer
}

func newHealthCmd(out io.Writer) *cobra.Command {
	healthCmd := &healthCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "health",
		Short: "Checl health about the Trinquet server",
		Long:  `All software has healths. This is Trinquet's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			return printHealth(client, healthCmd.out)
		},
	}
	return cmd
}

func printHealth(gRPCClient *utils.GRPCClient, out io.Writer) error {
	glog.V(1).Infof("Check health")

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := health.NewHealthServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}

	resp, err := client.Status(ctx, &health.StatusRequest{})
	if err != nil {
		return err
	}
	return printServerHealth(out, resp)
}

func printServerHealth(out io.Writer, resp *health.StatusResponse) error {

	table := tablewriter.NewWriter(out)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Service", "Status", "Text"})

	for _, service := range resp.Services {
		table.Append([]string{
			service.Name,
			service.Status,
			service.Text,
		})
	}
	table.Render()
	return nil
}
