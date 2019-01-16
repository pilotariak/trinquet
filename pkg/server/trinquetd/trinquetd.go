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
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/pilotariak/paleta/pkg/leagues"
	_ "github.com/pilotariak/paleta/pkg/leagues/ctpb"
	_ "github.com/pilotariak/paleta/pkg/leagues/ffpb"
	_ "github.com/pilotariak/paleta/pkg/leagues/lbpb"
	_ "github.com/pilotariak/paleta/pkg/leagues/lcapb"
	_ "github.com/pilotariak/paleta/pkg/leagues/lidfpb"
	"github.com/pilotariak/trinquet/pb/v1beta"
	"github.com/pilotariak/trinquet/pkg/api"
	"github.com/pilotariak/trinquet/pkg/config"
	"github.com/pilotariak/trinquet/pkg/storage"
	_ "github.com/pilotariak/trinquet/pkg/storage/boltdb"
)

func getStorage(conf *config.Configuration) (storage.Backend, error) {
	log.Info().Msgf("Create the backend using: %s", conf.Backend)
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
	log.Info().Str("backend", db.Name()).Msg("Initialize database")
	availablesLeagues := leagues.ListLeagues()
	for _, name := range availablesLeagues {
		log.Debug().Str("league", name).Msg("Manage league")
		leagueInfo, err := leagues.New(name)
		if err != nil {
			log.Error().Err(err).Msgf("Can't retrieve league: %s", name)
			break
		}
		leagueLevels := leagueInfo.Levels()
		var levels []*v1beta.Level
		for k, v := range leagueLevels {
			log.Debug().Str("league", name).Msgf("Add level for league %s %s", k, v)
			levels = append(levels, &v1beta.Level{
				Id:    k,
				Title: v,
			})
		}
		leagueDisciplines := leagueInfo.Disciplines()
		var disciplines []*v1beta.Discipline
		for k, v := range leagueDisciplines {
			log.Debug().Str("league", name).Msgf("Add discipline for league: %s %s", k, v)
			disciplines = append(disciplines, &v1beta.Discipline{
				Id:    k,
				Title: v,
			})
		}
		details := map[string]string{}
		for k, v := range leagueInfo.Details() {
			details[k] = v
		}
		league := &v1beta.League{
			Name:        name,
			Details:     details,
			Levels:      levels,
			Disciplines: disciplines,
		}
		log.Debug().Str("league", name).Msg("Save league informations")
		storage.StoreLeague(db, league)
	}
}

func StartServer(configFilename string) {

	conf, err := config.LoadFileConfig(configFilename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	serverAuth, err := newServerAuthentication(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create authentication")
	}

	db, err := getStorage(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to retrieve storage")
	}
	initializePelotaDatabase(db)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := fmt.Sprintf(":%d", conf.API.GrpcPort)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen for gRPC")
	}
	log.Info().Str("address", grpcAddr).Msg("gRPC address is listened")

	grpcServer, healthService, err := registerServer(db, serverAuth, conf, grpcAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to register gRPC server")
	}

	gwmux, err := registerGateway(ctx, fmt.Sprintf("localhost:%d", conf.API.GrpcPort))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to register JSON gateway")
	}

	httpmux := http.NewServeMux()
	httpmux.Handle("/v1/", gwmux)
	httpmux.Handle("/metrics", prometheus.Handler())
	httpmux.HandleFunc("/health", healthService.Handler)
	httpmux.HandleFunc("/version", api.VersionHandler)
	api.ServeStaticFile(httpmux)
	api.ServeSwagger(httpmux, "/swagger-ui/")

	// httpmux := http.NewServeMux()
	// httpmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`<html>
	//          <head><title>Trinquet</title></head>
	//          <body>
	//          <h1>Trinquet</h1>
	//          </body>
	//          </html>`))
	// })
	// httpmux.Handle("/v1beta/", gwmux)
	// httpmux.Handle("/metrics", prometheus.Handler())
	// httpmux.HandleFunc("/version", api.VersionHandler)
	// api.ServeSwagger(httpmux)

	log.Info().Str("address", grpcAddr).Msg("Start gRPC server")
	go grpcServer.Serve(lis)

	gwAddr := fmt.Sprintf(":%d", conf.API.RestPort)
	srv := &http.Server{
		Addr:    gwAddr,
		Handler: grpcHandlerFunc(grpcServer, httpmux),
	}
	log.Info().Str("address", gwAddr).Msg("Start HTTP server")
	srv.ListenAndServe()

}
