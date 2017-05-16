// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

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
