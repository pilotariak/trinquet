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
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/pilotariak/trinquet/api"
	"github.com/pilotariak/trinquet/version"
)

const (
	BANNER = "Trinquet"
)

var (
	debug bool
	vrsn  bool
	port  string
)

func init() {
	// parse flags
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "HTTP port for the server")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("%s v%s", BANNER, version.Version))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("%s v%s\n", BANNER, version.Version)
		os.Exit(0)
	}

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

}

func main() {
	logrus.Infof("[trinquet] Start web service")
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HelloHandler)
	mux.HandleFunc("/version", api.VersionHandler)
	mux.HandleFunc("/healthz", api.HealthzHandler)

	logrus.Debugf("[trinquet] Listent on :%s", port)
	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        api.LoggingHandler(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errChan := make(chan error, 10)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				logrus.Errorf("Error channel: %s", err.Error())
				os.Exit(0)
			}
		case s := <-signalChan:
			logrus.Errorf("Captured %v. Exiting...", s)
			os.Exit(0)
		}
	}

}
