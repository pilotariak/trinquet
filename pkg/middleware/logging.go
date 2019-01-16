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

package middleware

import (
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const gRPC = "gRPC"

// ServerLoggingInterceptor logs gRPC requests, errors and latency.
func ServerLoggingInterceptor(logSuccess bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		begin := time.Now()
		resp, err := handler(ctx, req)
		if err != nil {
			log.Error().Err(err).Msgf("%s %s (%v) %s", gRPC, info.FullMethod, err, time.Since(begin))
		} else if logSuccess {
			log.Info().Msgf("%s %s (success) %s", gRPC, info.FullMethod, time.Since(begin))
		}
		return resp, err
	}
}