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
	goflag "flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
	// init glog to get its flags
	_ "github.com/golang/glog"

	"github.com/pilotariak/trinquet/cmd/utils"
)

var (
	cliName           = "trinquetadm"
	helpMessage       = "Trinquetadm - The admin CLI for Trinquet"
	completionExample = `
               # Load the trinquetadm completion code for bash into the current shell
               source <(trinquetadm completion bash)

               # Write bash completion code to a file and source if from .bash_profile
               trinquetadm completion bash > ~/.trinquet/completion.bash.inc
               printf "\n# Trinquetadm shell completion\nsource '$HOME/.trinquet/completion.bash.inc'\n" >> $HOME/.bash_profile
               source $HOME/.bash_profile

               # Load the trinquetadm completion code for zsh[1] into the current shell
               source <(trinquetadm completion zsh)`
)

func init() {
	// Tell gRPC not to log to console.
	grpclog.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags))
}

// NewTrinquetadmCommand creates the `trinquetadm` command and its nested children.
func NewTrinquetadmCommand(out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   cliName,
		Short: "command-line tool to manage a Trinquet server",
		Long:  `Trinquetadm is a command-line tool to manage a Trinquet server.`,
	}
	rootCmd.AddCommand(
		newHealthCmd(out),
		newInfoCmd(out),
		utils.NewVersionCmd(out, helpMessage),
		utils.NewCompletionCommand(out, completionExample),
	)
	cobra.EnablePrefixMatching = true
	// add glog flags
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	// https://github.com/kubernetes/dns/pull/27/files
	goflag.CommandLine.Parse([]string{})

	return rootCmd
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := NewTrinquetadmCommand(os.Stdout)
	if err := cmd.Execute(); err != nil {
		fmt.Println(utils.RedOut(err))
		os.Exit(1)
	}
}
