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

package leagues

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// type Discipline struct {
// 	ID   string
// 	Name string
// }

// type Level struct {
// 	ID   string
// 	Name string
// }

type League interface {

	// Describe display informations
	//Describe()

	// Details send informations about the league
	Details() map[string]string

	// Display will print results on stdout
	Display(disciplineID string, levelID string) error

	// Levels retrieve available levels
	Levels() map[string]string // ([]Level, error)

	// Disciplines retrieve available pelota disciplines
	Disciplines() map[string]string // ([]Discipline, error)
}

type LeagueFunc func() (League, error)

var registeredLeagues = map[string](LeagueFunc){}

func RegisterLeague(name string, f LeagueFunc) {
	registeredLeagues[name] = f
}

func New(name string) (League, error) {
	f, ok := registeredLeagues[name]
	if !ok {
		return nil, fmt.Errorf("unknown league: %s", name)
	}
	return f()
}

func ListLeagues() []string {
	leagues := make([]string, 0, len(registeredLeagues))
	for name := range registeredLeagues {
		leagues = append(leagues, name)
	}
	sort.Strings(leagues)
	return leagues
}

func Describe(league League) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	for k, v := range league.Details() {
		table.Append([]string{k, v})
	}
	table.Render()
}
