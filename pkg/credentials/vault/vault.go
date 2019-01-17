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

package vault

import (
	"bytes"
	"crypto/tls"
	// "encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/vault/api"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
	"github.com/pilotariak/trinquet/pkg/transport"
)

const (
	name = "vault"

	namespace = "/secret/pilotariak"

	authretries       = 10
	retrydelayseconds = 3

	InvalidAPIKey         = "API key invalid"
	NoAuthenticationName  = "Authentication name is not set"
	NoAuthenticationRoles = "Authentication roles are not set"
)

type vaultSystem struct {
	Client   *api.Client
	Token    string
	RoleID   string
	SecretID string
}

func init() {
	credentials.RegisterCredentials(name, newVaultSystem)
}

func newVaultSystem(conf *config.Configuration) (credentials.CredentialsSystem, error) {
	log.Info().Str("credentials", name).Msgf("Configure Vault using: %s", conf.Credentials.Vault)

	if len(conf.Credentials.Vault.Address) == 0 {
		return nil, fmt.Errorf("Invalid Vault host: %s", conf.Credentials.Vault)
	}
	cfg := &api.Config{
		Address:    conf.Credentials.Vault.Address,
		HttpClient: cleanhttp.DefaultClient(),
		MaxRetries: 3,
	}

	cfg.HttpClient.Timeout = time.Second * 60
	httpTransport := cfg.HttpClient.Transport.(*http.Transport)
	httpTransport.TLSHandshakeTimeout = 10 * time.Second
	httpTransport.TLSClientConfig = &tls.Config{
		// MinVersion: tls.VersionTLS12,
		InsecureSkipVerify: true,
	}
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &vaultSystem{
		Client:   client,
		RoleID:   conf.Credentials.Vault.Roleid,
		SecretID: conf.Credentials.Vault.Secretid,
	}, nil
}

func (vs vaultSystem) Name() string {
	return name
}

func (vs vaultSystem) Authenticate(ctx context.Context, creds credentials.Credentials) error {
	log.Info().Str("credentials", name).Msgf("Check Vault credentials: %s", creds)

	if err := vs.login(); err != nil {
		return err
	}

	cuid, err := transport.ExtractMetadata(ctx, transport.UserID)
	if err != nil {
		return err
	}

	entry := fmt.Sprintf("%s/%s", namespace, cuid)
	log.Info().Str("credentials", name).Msgf("Check authentication for %s", entry)
	secret, err := vs.Client.Logical().Read(entry)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	if secret == nil || secret.Data == nil {
		return status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid entry: %s", entry))
	}
	log.Info().Str("credentials", name).Msgf("Vault data: %#v", secret.Data)
	apiKeyAuth, err := getEntry(secret, "apikey")
	if err != nil {
		return status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid authentication with this secret. %s %s", cuid, secret.Data))
	}
	if apiKeyAuth != creds.APIKey {
		return status.Errorf(codes.Unauthenticated, InvalidAPIKey)
	}
	name, err := getEntry(secret, "name")
	if err != nil {
		return status.Errorf(codes.Unauthenticated, NoAuthenticationName)
	}
	roles, err := getEntries(secret, "roles")
	if err != nil {
		return status.Errorf(codes.Unauthenticated, NoAuthenticationRoles)
	}
	log.Info().Str("credentials", name).Msgf("User correctly authenticated %s %s", cuid, roles)
	// headers := map[string]string{}
	// headers[transport.Username] = name
	// headers[transport.UserRoles] = roles
	// headers[transport.UserID] = cuid
	// return headers, nil
	return nil
}

func (vs vaultSystem) login() error {
	log.Info().Str("credentials", name).Msgf("Login into Vault")
	if len(vs.RoleID) == 0 ||
		len(vs.SecretID) == 0 {
		return status.Errorf(codes.Internal, "Invalid Vault configuration")
	}

	loginRequest := map[string]interface{}{
		"role_id":   vs.RoleID,
		"secret_id": vs.SecretID,
	}
	log.Info().Str("credentials", name).Msgf("Retrieve token")
	secretLogin, err := vs.Client.Logical().Write("auth/approle/login", loginRequest)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	vs.Client.SetToken(secretLogin.Auth.ClientToken)
	return nil
}

func getEntry(secret *api.Secret, key string) (string, error) {
	value, ok := secret.Data[key].(string)
	if !ok {
		return "", fmt.Errorf("Unknown entry")
	}
	return value, nil
}

func getEntries(secret *api.Secret, key string) (string, error) {
	var buffer bytes.Buffer
	values, ok := secret.Data[key].([]interface{})
	if !ok {
		return "", fmt.Errorf("Unknown entry")
	}
	for _, role := range values {
		log.Debug().Str("credentials", name).Msgf("Vault role: %s", role)
		buffer.WriteString(fmt.Sprintf("%s ", role))
	}
	return strings.TrimSpace(buffer.String()), nil
}
