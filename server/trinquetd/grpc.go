// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/mwitkow/go-grpc-middleware"
	"github.com/mwitkow/go-grpc-middleware/auth"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	ghealth "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/pilotariak/trinquet/api"
	_ "github.com/pilotariak/trinquet/auth/basic"
	_ "github.com/pilotariak/trinquet/auth/vault"
	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/middleware"
	"github.com/pilotariak/trinquet/pb/health"
	"github.com/pilotariak/trinquet/pb/info"
	"github.com/pilotariak/trinquet/pb/v1beta"
	"github.com/pilotariak/trinquet/storage"
)

func registerServer(backend storage.Backend, serverAuth *serverAuthentication, tracer opentracing.Tracer, conf *config.Configuration, grpcAddr string) (*grpc.Server, error) {
	glog.V(1).Info("Create the gRPC server")

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.ServerLoggingInterceptor(true),
				grpc_prometheus.UnaryServerInterceptor,
				otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
				grpc_auth.UnaryServerInterceptor(serverAuth.authenticate))),
	)

	services := []string{
		"LeagueService",
	}
	info.RegisterInfoServiceServer(server, api.NewInfoService(conf))
	healthService, err := api.NewHealthService(conf, grpcAddr, services)
	if err != nil {
		return nil, err
	}
	health.RegisterHealthServiceServer(server, healthService)

	v1beta.RegisterLeagueServiceServer(server, api.NewLeagueService(backend))

	healthServer := ghealth.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)
	healthServer.SetServingStatus("LeagueService", healthpb.HealthCheckResponse_SERVING)

	grpc_prometheus.Register(server)

	return server, nil
}

func registerGateway(ctx context.Context, addr string) (*runtime.ServeMux, error) {
	glog.V(1).Info("Create the REST gateway")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	gwmux := runtime.NewServeMux()
	if err := v1beta.RegisterLeagueServiceHandlerFromEndpoint(ctx, gwmux, addr, opts); err != nil {
		return nil, err
	}
	return gwmux, nil
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
