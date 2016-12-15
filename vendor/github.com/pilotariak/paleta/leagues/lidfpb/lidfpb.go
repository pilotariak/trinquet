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

package ctpb

import (
	// "fmt"
	// "os"

	"github.com/Sirupsen/logrus"
	// "github.com/olekukonko/tablewriter"

	"github.com/pilotariak/paleta/leagues"
)

const (
	uri = "http://lidfpb.euskalpilota.fr/resultats.php"
)

var (
	current = "20160401"

	disciplines = map[string]string{
		"2": "Trinquet / P.G. Pleine Masculin",
		"4": "Trinquet / P.G. Pleine Feminine",
		"6": "Trinquet / Xare",
	}

	levels = map[string]string{
		"1": "1ère Série - 1.Maila",
	}
)

func init() {
	leagues.RegisterLeague("lidfpb", newLIDFPBLeague)
}

type lidfpbLeague struct {
	Website string
	Name    string
}

func newLIDFPBLeague() (leagues.League, error) {
	return &lidfpbLeague{
		Website: "http://www.lidfpb.fr/",
		Name:    "Ligue d’île de France de pelote basque",
	}, nil
}

// func (l *lidfpbLeague) Describe() {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetRowLine(true)
// 	table.SetAutoWrapText(false)
// 	table.Append([]string{"Name", l.Name})
// 	table.Append([]string{"Website", l.Website})
// 	table.Render()
// }

func (l *lidfpbLeague) Details() map[string]string {
	return map[string]string{
		"Name":    l.Name,
		"Website": l.Website,
	}
}

func (l *lidfpbLeague) Levels() map[string]string {
	return levels
}

func (l *lidfpbLeague) Disciplines() map[string]string {
	return disciplines
}

func (l *lidfpbLeague) Display(disciplineID string, levelID string) error {
	logrus.Debugf("[lidfpb] Search results for %s %s", disciplineID, levelID)
	return leagues.Display(uri, disciplineID, levelID, current)
}
