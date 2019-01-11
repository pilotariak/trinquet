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

package basic

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pkg/auth"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/transport"
)

const (
	label = "BasicAuth"

	key = "basic"
)

type basicAuthSystem struct{}

func init() {
	auth.RegisterAuthentication(label, newBasicAuthSystem)
}

func newBasicAuthSystem(config *config.Configuration) (auth.Authentication, error) {
	return &basicAuthSystem{}, nil
}

func (ba basicAuthSystem) Name() string {
	return label
}

func (ba basicAuthSystem) Key() string {
	return key
}

func (ba basicAuthSystem) Credentials(ctx context.Context, username string, password string) (string, error) {
	log.Info().Str("auth", label).Msgf("Set credentials: %s", username)
	auth := username + ":" + password
	token := base64.StdEncoding.EncodeToString([]byte(auth))
	return token, nil
}

func (ba basicAuthSystem) Authenticate(ctx context.Context, token string) (map[string]string, error) {
	log.Info().Str("auth", label).Msgf("Check BasicAuth token: %s", token)
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("Can't check authentication: %s", err)
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return nil, fmt.Errorf("Not Authorized")
	}
	log.Info().Str("auth", label).Msgf("Auth: %s / %s", pair[0], pair[1])
	if pair[0] != "health" {
		if pair[0] != auth.Username || pair[1] != auth.Password {
			return nil, fmt.Errorf("Unauthorized")
		}
	}

	headers := map[string]string{}
	headers[transport.Username] = pair[0]
	return headers, nil
}
