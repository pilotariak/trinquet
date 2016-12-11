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
	"github.com/fatih/color"
	"google.golang.org/grpc"

	"github.com/pilotariak/trinquet/pb"
)

var (
	greenOut  = color.New(color.FgGreen).SprintFunc()
	yellowOut = color.New(color.FgYellow).SprintFunc()
	redOut    = color.New(color.FgRed).SprintFunc()
)

func getClient(uri string) (pb.LeagueServiceClient, error) {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	return pb.NewLeagueServiceClient(conn), nil
}
