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

package auth

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
)

type AuthenticationFunc func(config *config.Configuration) (Authentication, error)

var registeredSystems = map[string](AuthenticationFunc){}

func RegisterAuthentication(name string, f AuthenticationFunc) {
	registeredSystems[name] = f
}

// Authentication define an authentication system
type Authentication interface {

	// Name identify the system
	Name() string

	// Encode build username and password into a token
	Encode(ctx context.Context, username string, password string) (string, error)

	// Decode decode the token
	Decode(ctx context.Context, credentials credentials.Credentials, token string) (map[string]string, error)
}

// New returns a new authentication system using the name
func New(conf *config.Configuration) (Authentication, error) {
	log.Info().Msgf("Authentication setup: %s", conf.Auth)
	if conf.Auth == nil {
		return nil, fmt.Errorf("Invalid authentication configuration: %s", conf)
	}
	log.Debug().Msgf("Available systems: %s", registeredSystems)
	f, ok := registeredSystems[conf.Auth.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported authentication system: %s", conf.Auth.Name)
	}
	system, err := f(conf)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func GetAuthenticationHeader(authentication Authentication, token string) string {
	return fmt.Sprintf("%s %s", authentication.Name(), token)
}
