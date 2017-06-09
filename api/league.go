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

package api

import (
	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/messaging"
	"github.com/pilotariak/trinquet/pb/v1beta"
	"github.com/pilotariak/trinquet/storage"
	"github.com/pilotariak/trinquet/tracing"
)

const (
	LeagueServiceName = "LeagueService"
)

type LeagueService struct {
	Backend storage.Backend
}

func NewLeagueService(backend storage.Backend) *LeagueService {
	glog.V(2).Infof("Create the League service using %v", backend)
	service := &LeagueService{
		Backend: backend,
	}
	service.Register()
	return service
}

func (service *LeagueService) Register() {
	Services = append(Services, LeagueServiceName)
}

func (ls *LeagueService) List(ctx context.Context, request *v1beta.GetLeaguesRequest) (*v1beta.GetLeaguesResponse, error) {
	glog.V(1).Info("[league] List all leagues")
	span := tracing.GetParentSpan(ctx, messaging.ListLeaguesEvent)
	defer span.Finish()

	theleagues, err := storage.ListAll(ls.Backend)
	if err != nil {
		span.LogFields(log.Error(err))
		return nil, err
	}
	span.LogFields(log.Object("storage response", theleagues))
	return &v1beta.GetLeaguesResponse{Leagues: theleagues}, nil
}

func (ls *LeagueService) Create(ctx context.Context, request *v1beta.CreateLeagueRequest) (*v1beta.CreateLeagueResponse, error) {
	glog.V(1).Info("[league] Create a new league")
	return &v1beta.CreateLeagueResponse{}, nil
}

func (ls *LeagueService) Get(ctx context.Context, request *v1beta.GetLeagueRequest) (*v1beta.GetLeagueResponse, error) {
	glog.V(1).Info("[league] Retrieve a league")
	league, err := storage.RetrieveLeague(ls.Backend, request.Name)
	if err != nil {
		return nil, err
	}
	return &v1beta.GetLeagueResponse{
		League: league,
	}, nil
}
