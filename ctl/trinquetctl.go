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

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/ctl/commands"
)

var (
	// verbose       bool
	username      string
	password      string
	serverAddress string

	logging bool
	logFile = "/tmp/trinquetctl.log"

	// TrinquetctlCmd is the Trinquet cli
	TrinquetctlCmd = &cobra.Command{
		Use:   "trinquetctl",
		Short: "Trinquetctl is a CLI to use the Trinquet server",
	}
)

func init() {
	TrinquetctlCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username to use")
	TrinquetctlCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password to use")
	TrinquetctlCmd.PersistentFlags().StringVarP(&serverAddress, "serverAddr", "s", "localhost:8080", "The server address in the format of host:port")

	// TrinquetctlCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	// TrinquetctlCmd.PersistentFlags().BoolVar(&logging, "log", false, "Enable Logging")
	// TrinquetctlCmd.PersistentFlags().StringVar(&logFile, "logFile", "", "Log File path (if set, logging enabled automatically)")
}

func main() {
	addCommands()
	if err := TrinquetctlCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func addCommands() {
	TrinquetctlCmd.AddCommand(commands.VersionCmd)
	TrinquetctlCmd.AddCommand(commands.NewCmdLeagues())
	TrinquetctlCmd.AddCommand(commands.NewCmdLeague())
}
