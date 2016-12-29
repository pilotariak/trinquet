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

package jaeger

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"

	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/tracing"
)

const (
	label = "jaeger"
)

func init() {
	tracing.RegisterTracer(label, newTracer)
}

func newTracer(conf *config.Configuration) (opentracing.Tracer, error) {
	glog.V(1).Infof("Create OpenTracing tracer using Jaeger: %s %d", conf.Tracing.Zipkin.Host, conf.Tracing.Zipkin.Port)

	// Jaeger tracer can be initialized with a transport that will
	// report tracing Spans to a Zipkin backend
	transport, err := zipkin.NewHTTPTransport(
		fmt.Sprintf("%s:%d/api/v1/spans", conf.Tracing.Jaeger.Host, conf.Tracing.Jaeger.Port),
		zipkin.HTTPLogger(jaeger.StdLogger),
	)
	if err != nil {
		return nil, err
	}

	tracer, _ := jaeger.NewTracer(
		tracing.ServiceName,
		jaeger.NewConstSampler(true), // sample all traces
		jaeger.NewRemoteReporter(transport, nil),
	)
	// defer closer.Close()
	return tracer, nil
}
