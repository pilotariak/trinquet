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

package credentials

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pkg/config"
)

type Credentials struct {
	Username string
	Password string
	APIKey   string
}

type CredentialsFunc func(config *config.Configuration) (CredentialsSystem, error)

var registeredSystems = map[string](CredentialsFunc){}

func RegisterCredentials(name string, f CredentialsFunc) {
	registeredSystems[name] = f
}

// Credentials define a transport for credentials
type CredentialsSystem interface {

	// Name identify the system
	Name() string

	Authenticate(ctx context.Context, credentials Credentials) error
}

// New returns a new credentials system using the name
func New(conf *config.Configuration) (CredentialsSystem, error) {
	log.Info().Msgf("Credentials setup: %s", conf.Credentials)
	if conf.Credentials == nil {
		return nil, fmt.Errorf("Invalid credentials configuration: %s", conf)
	}
	log.Debug().Msgf("Available systems: %s", registeredSystems)
	f, ok := registeredSystems[conf.Credentials.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported credentials system: %s", conf.Credentials.Name)
	}
	system, err := f(conf)
	if err != nil {
		return nil, err
	}
	return system, nil
}
