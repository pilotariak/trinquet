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

package api

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/pilotariak/paleta/leagues"
	_ "github.com/pilotariak/paleta/leagues/ctpb"
	_ "github.com/pilotariak/paleta/leagues/ffpb"
	_ "github.com/pilotariak/paleta/leagues/lbpb"
	_ "github.com/pilotariak/paleta/leagues/lcapb"
	_ "github.com/pilotariak/paleta/leagues/lidfpb"

	"github.com/pilotariak/trinquet/pb"
)

type LeagueService struct {
}

func NewLeagueService() *LeagueService {
	glog.V(2).Info("Create the League service")
	return &LeagueService{}
}

func (ls *LeagueService) List(context.Context, *pb.GetLeaguesRequest) (*pb.GetLeaguesResponse, error) {
	glog.V(1).Info("[league] List all leagues")
	availablesLeagues := leagues.ListLeagues()
	size := len(availablesLeagues)
	glog.V(1).Infof("[league] Available leagues : %d", size)
	theleagues := make([]*pb.League, size)
	for i, name := range availablesLeagues {
		glog.V(2).Infof("[league] %s", name)
		theleagues[i] = &pb.League{
			Name:    name,
			Website: ""}
	}
	glog.V(1).Infof("[league] Response: %d %s", len(theleagues), theleagues)
	return &pb.GetLeaguesResponse{Leagues: theleagues}, nil
}

func (ls *LeagueService) Create(context.Context, *pb.CreateLeagueRequest) (*pb.CreateLeagueResponse, error) {
	glog.V(1).Info("[league] Create a new league")
	return &pb.CreateLeagueResponse{}, nil
}

func (ls *LeagueService) Get(context.Context, *pb.GetLeagueRequest) (*pb.GetLeagueResponse, error) {
	glog.V(1).Info("[league] Retrieve a league")
	return &pb.GetLeagueResponse{}, nil
}
