// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pilotariak/trinquet/pb"
	"github.com/pilotariak/trinquet/services"
)

const (
	port = 8080
)

func registerServer() *grpc.Server {
	glog.V(2).Info("[trinquet] Create the gRPC server")
	server := grpc.NewServer()
	pb.RegisterLeagueServiceServer(server, services.NewLeagueService())
	return server
}

func registerGateway(ctx context.Context) (*runtime.ServeMux, error) {
	glog.V(2).Info("[trinquet] Create the REST gateway")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	gwmux := runtime.NewServeMux()
	addr := fmt.Sprintf("localhost:%d", port)
	if err := pb.RegisterLeagueServiceHandlerFromEndpoint(ctx, gwmux, addr, opts); err != nil {
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

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.V(2).Info("preflight request for %s", r.URL.Path)
	return
}

func main() {
	glog.Infoln("[trinquet] Create the gRPC servers")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	glog.Infof("[trinquet] Listen on %s", grpcAddr)

	grpcServer := registerServer()
	gwmux, err := registerGateway(ctx)
	if err != nil {
		glog.Fatalf("Failed to register JSON gateway: %s", err.Error())
	}

	// httpmux := http.NewServeMux()
	// // httpmux.Handle("/", handler)
	// httpmux.Handle("/v1/", gwmux)

	// logrus.Infof("[trinquet] Listen on %s", addr)
	// http.ListenAndServe(addr, allowCORS(httpmux))
	glog.Infof("[trinquet] Start gRPC server on %s", grpcAddr)
	go grpcServer.Serve(lis)

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", 9090),
		Handler: grpcHandlerFunc(grpcServer, gwmux),
	}

	glog.Fatal(srv.ListenAndServe())

}
