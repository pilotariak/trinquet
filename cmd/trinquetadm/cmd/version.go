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
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/version"
)

type versionCmd struct {
	out io.Writer
}

func newVersionCmd(out io.Writer) *cobra.Command {
	versionCmd := &versionCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Trinquetctl",
		Long:  `All software has versions. This is Trinquet's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return versionCmd.printDiabloVersion()
		},
	}
	return cmd
}

func (cmd versionCmd) printDiabloVersion() error {
	fmt.Fprintf(cmd.out, "Trinquetctl - The CLI for Trinquet. v%s\n", version.Version)
	return nil
}
