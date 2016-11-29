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
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pb"
)

type LeagueService struct {
}

func newLeagueService() *LeagueService {
	return &LeagueService{}
}

func (ls *LeagueService) GetLeagues(context.Context, *pb.GetLeaguesRequest) (*pb.GetLeaguesResponse, error) {
	return &pb.GetLeaguesResponse{}, nil
}

func (ls *LeagueService) CreateLeague(context.Context, *pb.CreateLeagueRequest) (*pb.CreateLeagueResponse, error) {
	return &pb.CreateLeagueResponse{}, nil
}

func (ls *LeagueService) GetLeague(context.Context, *pb.GetLeagueRequest) (*pb.GetLeagueResponse, error) {
	return &pb.GetLeagueResponse{}, nil
}
