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

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/pilotariak/paleta/leagues"
	_ "github.com/pilotariak/paleta/leagues/ctpb"
	_ "github.com/pilotariak/paleta/leagues/ffpb"
	_ "github.com/pilotariak/paleta/leagues/lbpb"
	_ "github.com/pilotariak/paleta/leagues/lcapb"
	_ "github.com/pilotariak/paleta/leagues/lidfpb"
	"github.com/pilotariak/trinquet/api"
	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/middleware"
	"github.com/pilotariak/trinquet/pb"
	"github.com/pilotariak/trinquet/storage"
	_ "github.com/pilotariak/trinquet/storage/boltdb"
	"github.com/pilotariak/trinquet/tracing"
	_ "github.com/pilotariak/trinquet/tracing/jaeger"
	_ "github.com/pilotariak/trinquet/tracing/zipkin"
	"github.com/pilotariak/trinquet/version"
)

const (
	port = 8080
)

var (
	debug           bool
	vrs             bool
	defaultConfFile string
	// grpcPort int
	// gwPort   int
)

func registerServer(backend storage.Backend, tracer opentracing.Tracer) *grpc.Server {
	glog.V(1).Info("Create the gRPC server")

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.ServerLoggingInterceptor(true),
				grpc_prometheus.UnaryServerInterceptor,
				otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
				grpc_auth.UnaryServerInterceptor(authenticate))),
	)
	pb.RegisterLeagueServiceServer(server, api.NewLeagueService(backend))

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)
	healthServer.SetServingStatus("LeagueService", healthpb.HealthCheckResponse_SERVING)

	grpc_prometheus.Register(server)

	return server
}

func registerGateway(ctx context.Context, addr string) (*runtime.ServeMux, error) {
	glog.V(1).Info("Create the REST gateway")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	gwmux := runtime.NewServeMux()
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

// allowCORS allows Cross Origin Resource Sharing from any origin.
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

func getStorage(conf *config.Configuration) (storage.Backend, error) {
	glog.V(0).Infof("Create the backend using: %s", conf.Backend)
	db, err := storage.New(conf)
	if err != nil {
		return nil, err
	}
	err = db.Create()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initializePelotaDatabase(db storage.Backend) {
	glog.V(1).Infof("Initialize database")
	availablesLeagues := leagues.ListLeagues()
	for _, name := range availablesLeagues {
		glog.V(2).Infof("[league] %s", name)
		leagueInfo, err := leagues.New(name)
		if err != nil {
			glog.Errorf("Can't retrieve league: %s", name)
			break
		}
		leagueLevels := leagueInfo.Levels()
		var levels []*pb.Level
		for k, v := range leagueLevels {
			glog.V(2).Infof("For league %s add level %s %s", name, k, v)
			levels = append(levels, &pb.Level{
				Id:    k,
				Title: v,
			})
		}
		leagueDisciplines := leagueInfo.Disciplines()
		var disciplines []*pb.Discipline
		for k, v := range leagueDisciplines {
			glog.V(2).Infof("For league %s add discipline %s %s", name, k, v)
			disciplines = append(disciplines, &pb.Discipline{
				Id:    k,
				Title: v,
			})
		}
		details := map[string]string{}
		for k, v := range leagueInfo.Details() {
			details[k] = v
		}
		league := &pb.League{
			Name:        name,
			Details:     details,
			Levels:      levels,
			Disciplines: disciplines,
		}
		glog.V(2).Infof("Save League: %s", league)
		storage.StoreLeague(db, league)
	}
}

func init() {
	// parse flags
	flag.BoolVar(&vrs, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&defaultConfFile, "config", "trinquet.toml", "Configuration file to used")
	// flag.IntVar(&grpcPort, "grpcPort", 8080, "gRPC port to use")
	// flag.IntVar(&gwPort, "gwPort", 9090, "REST gateway port to use")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Trinquet v%s\n", version.Version))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrs {
		fmt.Printf("%s\n", version.Version)
		os.Exit(0)
	}
}

func main() {

	conf, err := config.LoadFileConfig(defaultConfFile)
	if err != nil {
		glog.Fatalf("failed to load configuration: %v", err)
	}

	db, err := getStorage(conf)
	if err != nil {
		glog.Fatalf("failed to load configuration: %v", err)
	}
	glog.V(1).Infof("Backend used: %s", db.Name())
	initializePelotaDatabase(db)

	tracer, err := tracing.New(conf)
	if err != nil {
		glog.Fatalf("failed to initialize OpenTracing: %v", err)
	}

	glog.V(0).Infoln("Create the gRPC servers")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := fmt.Sprintf(":%d", conf.API.GrpcPort)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	glog.V(0).Infof("Listen on %s", grpcAddr)

	grpcServer := registerServer(db, tracer)

	gwmux, err := registerGateway(ctx, fmt.Sprintf("localhost:%d", conf.API.GrpcPort))
	if err != nil {
		glog.Fatalf("Failed to register JSON gateway: %s", err.Error())
	}

	httpmux := http.NewServeMux()
	httpmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Trinquet</title></head>
             <body>
             <h1>Trinquet</h1>
             </body>
             </html>`))
	})
	httpmux.HandleFunc("/version", api.VersionHandler)
	httpmux.Handle("/v1/", gwmux)
	httpmux.Handle("/metrics", prometheus.Handler())
	healthHandler, err := api.NewHealthHandler(grpcAddr)
	if err != nil {
		glog.Fatalf("Failed to create Health monitor: %s", err.Error())
	}
	httpmux.HandleFunc("/healthz", healthHandler.Handle)
	api.ServeSwagger(httpmux)

	glog.V(0).Infof("Start gRPC server on %s", grpcAddr)
	go grpcServer.Serve(lis)

	gwAddr := fmt.Sprintf(":%d", conf.API.RestPort)
	srv := &http.Server{
		Addr:    gwAddr,
		Handler: grpcHandlerFunc(grpcServer, httpmux),
	}
	glog.V(0).Infof("Start HTTP server on %s", gwAddr)
	glog.Fatal(srv.ListenAndServe())

}
