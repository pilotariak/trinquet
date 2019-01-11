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

package api

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pb/info"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/rbac"
	"github.com/pilotariak/trinquet/pkg/version"
)

const (
	infoServiceName = "health"
)

type InfoService struct {
	Version string
}

func NewInfoService(conf *config.Configuration) *InfoService {
	log.Info().Str("service", infoServiceName).Msg("Create the info service")
	rbac.AddRoles("info", "InfoService", map[string][]string{
		"Get": []string{rbac.AdminRole},
	})
	return &InfoService{
		Version: version.Version,
	}
}

type apiVersion struct {
	Version string `json:"version"`
}

func (service *InfoService) Get(ctx context.Context, req *info.GetInfoRequest) (*info.GetInfoResponse, error) {
	log.Info().Str("service", infoServiceName).Msg("Retrieve informations")

	resp := &info.GetInfoResponse{
		Version: service.Version,
	}
	log.Info().Str("service", infoServiceName).Msgf("Info response: %s", resp)
	return resp, nil
}
