// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package api

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"

	"github.com/pilotariak/trinquet/auth"
	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/messaging"
	"github.com/pilotariak/trinquet/pb/health"
	"github.com/pilotariak/trinquet/pkg/rbac"
	"github.com/pilotariak/trinquet/tracing"
	"github.com/pilotariak/trinquet/transport"
)

type HealthService struct {
	Authentication auth.Authentication
	HealthUser     string
	HealthKey      string
	URI            string
	Services       []string
}

func NewHealthService(conf *config.Configuration, uri string, services []string) (*HealthService, error) {
	glog.V(2).Info("Create the health service")
	rbac.AddRoles("health", "HealthService", map[string][]string{
		"Status": []string{rbac.AdminRole},
	})
	authentication, err := auth.New(conf)
	if err != nil {
		return nil, err
	}
	return &HealthService{
		// Conf:     conf,
		Authentication: authentication,
		HealthKey:      conf.Auth.Vault.HealthKey,
		HealthUser:     conf.Auth.Vault.HealthUser,
		URI:            uri,
		Services:       services,
	}, nil
}

func (service *HealthService) Status(ctx context.Context, req *health.StatusRequest) (*health.StatusResponse, error) {
	glog.V(1).Info("Check Health services")

	span := tracing.GetParentSpan(ctx, messaging.HealthEvent)
	defer span.Finish()

	conn, err := grpc.Dial(service.URI, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := healthpb.NewHealthClient(conn)
	token, err := service.Authentication.Credentials(ctx, span, service.HealthUser, service.HealthKey)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("Auth token: %s", token)
	md := metadata.New(map[string]string{
		transport.Authorization: auth.GetAuthenticationHeader(service.Authentication, token),
		transport.UserID:        service.HealthUser,
	})
	newCtx := metadata.NewContext(ctx, md)

	servicesStatus := []*health.ServiceStatus{}
	for _, service := range service.Services {
		glog.V(2).Infof("Check health service: %s", service)
		resp, err := client.Check(newCtx, &healthpb.HealthCheckRequest{
			Service: service,
		})
		if err != nil {
			servicesStatus = append(servicesStatus, &health.ServiceStatus{
				Name:   service,
				Status: "KO",
				Text:   err.Error(),
			})
		} else {
			servicesStatus = append(servicesStatus, &health.ServiceStatus{
				Name:   service,
				Status: "OK",
				Text:   fmt.Sprintf("%s", resp.Status),
			})
		}
	}

	resp := &health.StatusResponse{}
	resp.Services = servicesStatus

	glog.V(0).Infof("Health response: %s", resp)
	span.LogFields(log.Object("response", resp))
	return resp, nil
}
