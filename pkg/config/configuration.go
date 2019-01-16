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

package config

import (
	//"log"

	"github.com/BurntSushi/toml"
)

// Configuration holds configuration for Enigma.
type Configuration struct {
	Backend     string
	API         *APIConfiguration
	BoltDB      *BoltDBConfiguration
	Auth        *AuthConfiguration
	Credentials *CredentialsConfiguration
}

// New returns a Configuration with default values
func New() *Configuration {
	return &Configuration{
		Backend:     "boltdb",
		BoltDB:      &BoltDBConfiguration{},
		API:         &APIConfiguration{},
		Auth:        &AuthConfiguration{},
		Credentials: &CredentialsConfiguration{},
	}
}

// LoadFileConfig returns a Configuration from reading the specified file (a toml file).
func LoadFileConfig(file string) (*Configuration, error) {
	configuration := New()
	if _, err := toml.DecodeFile(file, configuration); err != nil {
		return nil, err
	}
	return configuration, nil
}

// APIConfiguration defines the configuration for the gRPC and REST api
type APIConfiguration struct {
	GrpcPort int
	RestPort int
}

// BoltDBConfiguration defines the configuration for BoltDB storage backend
type BoltDBConfiguration struct {
	Bucket string
	File   string
}

type CredentialsConfiguration struct {
	Name  string
	Text  *TextCredentialsConfiguration
	Vault *VaultCredentialsConfiguration
}

type TextCredentialsConfiguration struct {
	Username string
	Password string
}

type VaultCredentialsConfiguration struct {
	Address    string
	Roleid     string
	Secretid   string
	HealthUser string `toml:"healthuser"`
	HealthKey  string `toml:"healthkey"`
}

type AuthConfiguration struct {
	Name      string
	BasicAuth *BasicAuthConfiguration
	JWT       *JWTConfiguration
}

type BasicAuthConfiguration struct {
}

type JWTConfiguration struct {
	SigningKey string
	Secret     string
}
