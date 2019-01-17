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

package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/pkg/auth"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
	"github.com/pilotariak/trinquet/pkg/transport"
)

const (
	name   = "JWT"
	scheme = "bearer"
)

var (
	jwtTimeout = time.Minute * 5
)

type token struct {
	username string `json:"username"`
	jwt.StandardClaims
}

type jwtSystem struct {
	signingKey []byte
	secret     []byte
}

func init() {
	auth.RegisterAuthentication(name, newJwtSystem)
}

func newJwtSystem(conf *config.Configuration) (auth.Authentication, error) {
	log.Info().Str("auth", name).Msgf("Configuration: %s", conf.Auth.BasicAuth)
	return &jwtSystem{
		signingKey: []byte(conf.Auth.JWT.SigningKey),
		secret:     []byte(conf.Auth.JWT.Secret),
	}, nil
}

func (js jwtSystem) Name() string {
	return name
}

func (js jwtSystem) Scheme() string {
	return scheme
}

func (js jwtSystem) Encode(ctx context.Context, username string, password string) (string, error) {
	log.Info().Str("auth", name).Msgf("Encode credentials: %s", username)
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      now.Add(jwtTimeout).Unix(),
		"iat":      now.Unix(),
	})
	return token.SignedString(js.signingKey)
}

func (js jwtSystem) Decode(ctx context.Context, credsSystem credentials.CredentialsSystem, tokenString string) (map[string]string, error) {
	log.Info().Str("auth", name).Msgf("Decode token: %s", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return js.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		creds := credentials.Credentials{
			Username: fmt.Sprintf("%s", claims["username"]),
			Password: fmt.Sprintf("%s", claims["password"]),
		}
		if err := credsSystem.Authenticate(ctx, creds); err != nil {
			return nil, err
		}
		headers := map[string]string{}
		headers[transport.Username] = creds.Username
		return headers, nil

	}
	return nil, fmt.Errorf("Invalid JWT token")
}
