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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"

	"github.com/pilotariak/trinquet/pb/health"
	"github.com/pilotariak/trinquet/pkg/auth"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/rbac"
	"github.com/pilotariak/trinquet/pkg/transport"
)

const (
	healthServiceName = "health"
)

type HealthService struct {
	Authentication auth.Authentication
	HealthUser     string
	HealthKey      string
	URI            string
	Services       []string
}

func NewHealthService(conf *config.Configuration, uri string, services []string) (*HealthService, error) {
	log.Info().Str("service", healthServiceName).Msg("Create the service")
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
	log.Info().Msg("Check Health services")

	client, newCtx, err := service.getClient(ctx)
	if err != nil {
		return nil, err
	}

	servicesStatus := service.checkServices(client, newCtx)
	resp := &health.StatusResponse{
		Services: servicesStatus,
	}

	log.Info().Msgf("Health response: %s", resp)
	return resp, nil
}

func (service *HealthService) getClient(ctx context.Context) (healthpb.HealthClient, context.Context, error) {
	conn, err := grpc.Dial(service.URI, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	token, err := service.Authentication.Credentials(ctx, service.HealthUser, service.HealthKey)
	if err != nil {
		return nil, nil, err
	}
	log.Info().Str("token", token).Msg("Authicatation using token")
	md := metadata.New(map[string]string{
		transport.Authorization: auth.GetAuthenticationHeader(service.Authentication, token),
		transport.UserID:        service.HealthUser,
	})
	newCtx := metadata.NewIncomingContext(ctx, md)
	client := healthpb.NewHealthClient(conn)
	return client, newCtx, nil
}

func (service *HealthService) checkServices(client healthpb.HealthClient, ctx context.Context) []*health.ServiceStatus {
	servicesStatus := []*health.ServiceStatus{}
	for _, service := range service.Services {
		log.Info().Str("service", service).Msg("Check health service")
		resp, err := client.Check(ctx, &healthpb.HealthCheckRequest{
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
	return servicesStatus
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func sendErrorResponse(w http.ResponseWriter, err error) {
	response := errorResponse{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("%s", err),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (service *HealthService) Handler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Retrieve health")

	ctx := context.Background()
	client, newCtx, err := service.getClient(ctx)
	if err != nil {
		sendErrorResponse(w, err)
		return
	}

	servicesStatus := service.checkServices(client, newCtx)
	log.Info().Msgf("Health response: %s", servicesStatus)

	data, err := defaultMarshaler.Marshal(servicesStatus)
	if err != nil {
		sendErrorResponse(w, err)
		return
	}
	var response []*health.ServiceStatus
	if err := defaultMarshaler.Unmarshal(data, &response); err != nil {
		sendErrorResponse(w, err)
		return
	}

	json.NewEncoder(w).Encode(response)
	return
}

// func (service *HealthService) Status(ctx context.Context, req *health.StatusRequest) (*health.StatusResponse, error) {
// 	log.Info().Str("service", healthServiceName).Msg("Check status")

// 	conn, err := grpc.Dial(service.URI, grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := healthpb.NewHealthClient(conn)
// 	token, err := service.Authentication.Credentials(ctx, service.HealthUser, service.HealthKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Info().Str("service", healthServiceName).Msgf("Auth token: %s", token)
// 	md := metadata.New(map[string]string{
// 		transport.Authorization: auth.GetAuthenticationHeader(service.Authentication, token),
// 		transport.UserID:        service.HealthUser,
// 	})
// 	newCtx := metadata.NewOutgoingContext(ctx, md)

// 	servicesStatus := []*health.ServiceStatus{}
// 	for _, service := range service.Services {
// 		log.Info().Str("service", healthServiceName).Msgf("Check health service: %s", service)
// 		resp, err := client.Check(newCtx, &healthpb.HealthCheckRequest{
// 			Service: service,
// 		})
// 		if err != nil {
// 			servicesStatus = append(servicesStatus, &health.ServiceStatus{
// 				Name:   service,
// 				Status: "KO",
// 				Text:   err.Error(),
// 			})
// 		} else {
// 			servicesStatus = append(servicesStatus, &health.ServiceStatus{
// 				Name:   service,
// 				Status: "OK",
// 				Text:   fmt.Sprintf("%s", resp.Status),
// 			})
// 		}
// 	}

// 	resp := &health.StatusResponse{}
// 	resp.Services = servicesStatus

// 	log.Info().Str("service", healthServiceName).Msgf("Health response: %s", resp)
// 	return resp, nil
// }
