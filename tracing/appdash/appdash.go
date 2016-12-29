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

package appdash

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/tracing"
)

const (
	tracerLabel = "appdash"
)

func init() {
	tracing.RegisterTracer(tracerLabel, newTracer)
}

func newTracer(conf *config.Configuration) (opentracing.Tracer, error) {
	glog.V(1).Infof("Create OpenTracing tracer using Appdash: %s %d", conf.Tracing.Appdash.Host, conf.Tracing.Appdash.Port)
	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(fmt.Sprintf("%s:%d", conf.Tracing.Appdash.Host, conf.Tracing.Appdash.Port)))
	return tracer, nil
}
