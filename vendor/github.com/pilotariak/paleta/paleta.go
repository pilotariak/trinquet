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
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"

	"github.com/pilotariak/paleta/leagues"
	_ "github.com/pilotariak/paleta/leagues/ctpb"
	_ "github.com/pilotariak/paleta/leagues/ffpb"
	_ "github.com/pilotariak/paleta/leagues/lbpb"
	_ "github.com/pilotariak/paleta/leagues/lcapb"
	_ "github.com/pilotariak/paleta/leagues/lidfpb"
	"github.com/pilotariak/paleta/version"
)

var (
	debug        bool
	vrsn         bool
	league       string
	listleagues  bool
	levels       bool
	levelID      int
	disciplines  bool
	disciplineID int
	describe     bool
)

func init() {
	// parse flags
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.BoolVar(&listleagues, "leagues", false, "Display available leagues")
	flag.StringVar(&league, "league", "", "League key")
	flag.BoolVar(&levels, "levels", false, "Display available levels for a league")
	flag.IntVar(&levelID, "level", -1, "Level ID to look on")
	flag.BoolVar(&disciplines, "disciplines", false, "Display available disciplines for a league")
	flag.IntVar(&disciplineID, "discipline", -1, "Discipline ID to look on")
	flag.BoolVar(&describe, "describe", false, "Describe league")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Paleta v%s", version.Version))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("Paleta v%s\n", version.Version)
		os.Exit(0)
	}

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if listleagues {
		fmt.Println("Leagues:")
		for _, name := range leagues.ListLeagues() {
			fmt.Printf("- %s\n", name)
		}
		os.Exit(0)
	}

	if len(league) == 0 {
		usageAndExit("League name can't be empty.", 1)
	}

}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

func main() {
	// On ^C, or SIGTERM handle exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			logrus.Infof("Received %s, exiting.", sig.String())
			os.Exit(0)
		}
	}()

	league, err := leagues.New(league)
	if err != nil {
		logrus.Errorf("Can't retrieve league: %s", err.Error())
		os.Exit(1)
	}

	if levels {
		fmt.Println("Levels:")
		for k, v := range league.Levels() {
			fmt.Printf("- [%s] %s\n", k, v)
		}
		os.Exit(0)
	}
	if disciplines {
		fmt.Println("Disciplines:")
		for k, v := range league.Disciplines() {
			fmt.Printf("- [%s] %s\n", k, v)
		}
		os.Exit(0)
	}
	if describe {
		leagues.Describe(league)
		os.Exit(0)
	}

	if levelID == -1 || disciplineID == -1 {
		usageAndExit("Please specify level and discipline", 1)
	}

	league.Display(fmt.Sprintf("%d", disciplineID), fmt.Sprintf("%d", levelID))
}
