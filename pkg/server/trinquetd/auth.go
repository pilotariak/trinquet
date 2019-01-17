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

package trinquetd

import (
	"github.com/mwitkow/go-grpc-middleware/auth"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/pilotariak/trinquet/pkg/auth"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/credentials"
)

type serverAuthentication struct {
	authentication auth.Authentication
	credentials    credentials.CredentialsSystem
}

func newServerAuthentication(conf *config.Configuration) (*serverAuthentication, error) {
	log.Info().Msgf("Create the server authentication system")
	authentication, err := auth.New(conf)
	if err != nil {
		return nil, err
	}
	creds, err := credentials.New(conf)
	if err != nil {
		return nil, err
	}
	return &serverAuthentication{
		authentication: authentication,
		credentials:    creds,
	}, nil
}

func (sa *serverAuthentication) authenticate(ctx context.Context) (context.Context, error) {
	log.Info().Msgf("Check authentication using %s", sa.authentication.Name())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Debug().Msgf("Metadata: %s", md)
	}

	log.Debug().Msgf("Extract informations from headers")
	token, err := grpc_auth.AuthFromMD(ctx, sa.authentication.Scheme())
	if err != nil {
		return nil, err
	}

	headers, err := sa.authentication.Decode(ctx, sa.credentials, token)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}
	log.Info().Msgf("Authentication add headers: %s", headers)
	return metadata.NewIncomingContext(ctx, metadata.New(headers)), nil
}
