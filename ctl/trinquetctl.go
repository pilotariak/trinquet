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

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pilotariak/trinquet/ctl/commands"
)

var (
	logging bool
	logFile = "/tmp/trinquetctl.log"

	// TrinquetctlCmd is the Trinquet cli
	TrinquetctlCmd = &cobra.Command{
		Use:   "trinquetctl",
		Short: "Trinquetctl is a CLI to use the Trinquet server",
	}
)

func init() {
	// TrinquetctlCmd.Flags().BoolVarP(&version, "version", "v", version, "Print version info and exit")
	// TrinquetctlCmd.Flags().StringP("http-address", "H", httpAddress, "Http address of the gRPC server")

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
}

// func main() {

// 	var (
// 		debug bool
// 		vrs   bool
// 		uri   string
// 	)

// 	// parse flags
// 	flag.BoolVar(&vrs, "version", false, "print version and exit")
// 	flag.BoolVar(&debug, "d", false, "run in debug mode")
// 	flag.StringVar(&uri, "uri", "localhost:8080", "URI of the server")

// 	flag.Usage = func() {
// 		fmt.Fprint(os.Stderr, fmt.Sprintf("Trinquet v%s\n", version.Version))
// 		flag.PrintDefaults()
// 	}

// 	flag.Parse()

// 	if vrs {
// 		fmt.Printf("%s\n", version.Version)
// 		os.Exit(0)
// 	}

// 	// Set up a connection to the gRPC server.
// 	conn, err := grpc.Dial(uri, grpc.WithInsecure())
// 	if err != nil {
// 		glog.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewLeagueServiceClient(conn)
// 	glog.Infoln("[trinquet] Retrieve all leagues")
// 	fmt.Println(greenOut("Availables leagues:"))
// 	resp, err := client.List(context.Background(), &pb.GetLeaguesRequest{})
// 	if err != nil {
// 		glog.Fatalf("[trinquet] Could not retrieve leagues: %v", err)
// 		fmt.Println(redOut(fmt.Sprintf("Could not retrieve leagues: %v", err.Error())))
// 		os.Exit(0)
// 	}
// 	glog.Infof("[trinquet] Available leagues: %s", resp)
// 	for _, league := range resp.Leagues {
// 		fmt.Printf("- %s\n", league.Name)
// 	}
// }
