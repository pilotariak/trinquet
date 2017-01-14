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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// HealthHandler define a service
type HealthHandler struct {
	HealthClient healthpb.HealthClient
}

func NewHealthHandler(uri string) (*HealthHandler, error) {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := healthpb.NewHealthClient(conn)
	return &HealthHandler{
		HealthClient: client,
	}, nil
}

// HealthResponse contains current health status.
type HealthResponse struct {
	Status   string            `json:"status"`
	Gateway  string            `json:"gateway"`
	Services map[string]string `json:"services"`
}

func (handler *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	glog.V(1).Infof("[Health] handler")
	response := HealthResponse{
		Services: map[string]string{},
	}
	resp, err := handler.HealthClient.Check(context.Background(), &healthpb.HealthCheckRequest{Service: "LeagueService"})
	if err != nil {
		response.Services["LeagueService"] = fmt.Sprintf("KO: %s", err)
	} else {
		response.Services["LeagueService"] = fmt.Sprintf("%s", resp.Status)
	}
	response.Status = "OK"
	response.Gateway = "OK"
	json.NewEncoder(w).Encode(response)
	return
}
