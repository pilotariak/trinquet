// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package jaeger

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

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
	glog.V(1).Infof("Create OpenTracing tracer using Jaeger: %s %d", conf.Tracing.Jaeger.Host, conf.Tracing.Jaeger.Port)

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", conf.Tracing.Jaeger.Host, conf.Tracing.Jaeger.Port),
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	// Initialize tracer with a logger and a metrics factory
	tracer, _, err := cfg.New(
		tracing.ServiceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	// defer closer.Close()
	return tracer, nil
}
