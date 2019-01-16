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

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pkg/auth"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
	// "github.com/pilotariak/trinquet/pkg/transport"
)

const (
	// label = "BasicAuth"
	key = "BasicAuth"
)

type textSystem struct {
}

func init() {
	auth.RegisterAuthentication(key, newBasicAuthSystem)
}

func newBasicAuthSystem(conf *config.Configuration) (auth.Authentication, error) {
	log.Info().Str("auth", key).Msgf("Configuration: %s", conf.Auth.BasicAuth)
	return &textSystem{}, nil
}

func (ba textSystem) Name() string {
	return key
}

func (ba textSystem) Encode(ctx context.Context, username string, password string) (string, error) {
	log.Info().Str("auth", key).Msgf("Set credentials: %s", username)
	auth := username + ":" + password
	token := base64.StdEncoding.EncodeToString([]byte(auth))
	return token, nil
}

func (ba textSystem) Decode(ctx context.Context, credentials credentials.Credentials, token string) (map[string]string, error) {
	log.Info().Str("auth", key).Msgf("Token to decode: %s", token)
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("Can't decode authentication token: %s", err)
	}

	if err := credentials.Authenticate(ctx, string(b)); err != nil {
		return nil, err
	}

	headers := map[string]string{}
	return headers, nil
}
