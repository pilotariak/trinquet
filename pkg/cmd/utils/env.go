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

package utils

import (
	"os"

	"github.com/rs/zerolog/log"
)

const (
	// UsernameEnvVar is the environment variable that points to the username
	UsernameEnvVar = "TRINQUET_USERNAME"

	// ApikeyEnvVar is the environment variable that points to the APIKEY
	ApikeyEnvVar = "TRINQUET_APIKEY"

	// GrpcAddr is the environment variable that points to the gRPC server
	GrpcAddr = "TRINQUET_SERVER"

	// AuthMethod is the system to used for authentication
	AuthMethod = "TRINQUET_AUTH_SYSTEM"
)

const (
	defaultAuthSystem = "basic"
)

var (
	username      string
	password      string
	authSystem    string
	serverAddress string
	// RestAddress   string
)

func setupFromEnvironmentVariables() {
	username = os.Getenv(UsernameEnvVar)
	password = os.Getenv(ApikeyEnvVar)
	serverAddress = os.Getenv(GrpcAddr)
	authSystem = os.Getenv(AuthMethod)
	if len(authSystem) == 0 {
		authSystem = defaultAuthSystem
	}
	log.Debug().Msgf("Env: %s %s %s", username, serverAddress, authSystem)
}