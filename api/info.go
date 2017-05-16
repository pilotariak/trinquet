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

	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/messaging"
	"github.com/pilotariak/trinquet/pb/info"
	"github.com/pilotariak/trinquet/pkg/rbac"
	"github.com/pilotariak/trinquet/tracing"
	"github.com/pilotariak/trinquet/version"
)

type InfoService struct {
	Version string
}

func NewInfoService(conf *config.Configuration) *InfoService {
	glog.V(2).Info("Create the info service")
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
	glog.V(1).Info("Retrieve informations")

	span := tracing.GetParentSpan(ctx, messaging.InfoEvent)
	defer span.Finish()

	resp := &info.GetInfoResponse{
		Version: service.Version,
	}
	glog.V(0).Infof("Info response: %s", resp)
	span.LogFields(log.Object("response", resp))
	return resp, nil
}
