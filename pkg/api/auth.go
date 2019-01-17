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

	"github.com/pilotariak/trinquet/pb/v1"
	"github.com/pilotariak/trinquet/pkg/auth"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
)

const (
	AuthServiceName = "AuthenticationService"
)

type AuthService struct {
	authentication auth.Authentication
	credentials    credentials.CredentialsSystem
}

func NewAuthenticationService(conf *config.Configuration) (*AuthService, error) {
	log.Info().Str("service", LeagueServiceName).Msgf("Create the authentication service using %v", conf)
	authentication, err := auth.New(conf)
	if err != nil {
		return nil, err
	}
	creds, err := credentials.New(conf)
	if err != nil {
		return nil, err
	}
	service := &AuthService{
		authentication: authentication,
		credentials:    creds,
	}
	service.Register()
	return service, nil
}

func (service *AuthService) Register() {
	Services = append(Services, LeagueServiceName)
}

func (service *AuthService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	log.Info().Str("service", AuthServiceName).Msgf("No authentication needed for this service: %s", fullMethodName)
	return ctx, nil
}

func (service *AuthService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	log.Info().Str("service", AuthServiceName).Msgf("Credentials received: %s", request)
	creds := credentials.Credentials{
		Username: request.Username,
		Password: request.Password,
	}
	if err := service.credentials.Authenticate(ctx, creds); err != nil {
		return nil, err
	}
	token, err := service.authentication.Encode(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	log.Info().Str("service", AuthServiceName).Msgf("Authentication is valid. Token :%s", token)
	return &v1.LoginResponse{
		Token: token,
	}, nil
}
