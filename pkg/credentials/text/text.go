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

package text

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
)

const (
	name = "text"
)

type textSystem struct {
	username string
	password string
}

func init() {
	credentials.RegisterCredentials(name, newTextSystem)
}

func newTextSystem(conf *config.Configuration) (credentials.Credentials, error) {
	log.Info().Str("credentials", name).Msgf("Configure Text using: %s", conf.Credentials.Text)
	return &textSystem{
		username: conf.Credentials.Text.Username,
		password: conf.Credentials.Text.Password,
	}, nil
}

func (ts textSystem) Name() string {
	return name
}

func (ts textSystem) Authenticate(ctx context.Context, token string) error {
	log.Info().Str("credentials", name).Msgf("Authenticate with token: %s", token)

	pair := strings.SplitN(token, ":", 2)
	if len(pair) != 2 {
		return fmt.Errorf("Not Authorized")
	}
	if pair[0] != ts.username || pair[1] != ts.password {
		return fmt.Errorf("Unauthorized")
	}
	log.Info().Str("credentials", name).Msgf("User correctly authenticated %s ", pair[0])
	return nil
}
