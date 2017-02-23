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

package tracing

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/config"
)

const (
	// ServiceName used to setup the tracer
	ServiceName string = "trinquet"
)

type TracerFunc func(conf *config.Configuration) (opentracing.Tracer, error)

var registeredTracers = map[string](TracerFunc){}

func RegisterTracer(name string, f TracerFunc) {
	registeredTracers[name] = f
}

func New(conf *config.Configuration) (opentracing.Tracer, error) {
	glog.V(1).Infof("Opentracing setup: %s", conf.Tracing)
	f, ok := registeredTracers[conf.Tracing.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported tracer: %s", conf.Tracing.Name)
	}
	tracer, err := f(conf)
	if err != nil {
		return nil, err
	}

	// explicitly set our tracer to be the default tracer.
	opentracing.SetGlobalTracer(tracer)
	return tracer, nil
}

func GetSpan(ctx context.Context, operationName string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		glog.V(2).Infof("Got parent span for service")
	} else {
		parentSpan, ctx = opentracing.StartSpanFromContext(ctx, operationName)
	}
	return parentSpan
}
