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
	"io"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/pkg/server/trinquetd"
)

var (
	config string
)

type runCmd struct {
	out io.Writer
}

func newRunCmd(out io.Writer) *cobra.Command {
	runCmd := &runCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a Trinquet server",
		Long:  `All software has infos. This is Trinquet's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(config) == 0 {
				return errors.New("missing configuration filename")
			}
			runServer(runCmd.out, config)
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&config, "config", "trinquet.toml", "Configuration filename")
	return cmd
}

func runServer(out io.Writer, config string) {
	log.Info().Msg("Start the server")
	trinquetd.StartServer(config)
}
