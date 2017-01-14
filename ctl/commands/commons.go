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

package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/pb"
	"github.com/pilotariak/trinquet/tracing"
	_ "github.com/pilotariak/trinquet/tracing/appdash"
	_ "github.com/pilotariak/trinquet/tracing/jaeger"
	_ "github.com/pilotariak/trinquet/tracing/zipkin"
)

var (
	greenOut  = color.New(color.FgGreen).SprintFunc()
	yellowOut = color.New(color.FgYellow).SprintFunc()
	redOut    = color.New(color.FgRed).SprintFunc()

	httpAddress string

	tracerName     string
	zipkinAddress  string
	zipkinPort     int
	appdashAddress string
	appdashPort    int
)

func createConfiguration() (*config.Configuration, error) {
	switch tracerName {
	case "zipkin":
		return &config.Configuration{
			Tracing: &config.TracingConfiguration{
				Name: tracerName,
				Zipkin: &config.ZipkinConfiguration{
					Host: zipkinAddress,
					Port: zipkinPort,
				},
			},
		}, nil
	default:
		return nil, fmt.Errorf("OpenTracing tracer could not be empty.")
	}
}

func getClient(uri string) (pb.LeagueServiceClient, opentracing.Tracer, error) {
	conf, err := createConfiguration()
	if err != nil {
		return nil, nil, err
	}

	tracer, err := tracing.New(conf)
	if err != nil {
		return nil, nil, err
	}

	conn, err := grpc.Dial(
		uri,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		return nil, nil, err
	}
	// defer conn.Close()

	return pb.NewLeagueServiceClient(conn), tracer, nil
}
