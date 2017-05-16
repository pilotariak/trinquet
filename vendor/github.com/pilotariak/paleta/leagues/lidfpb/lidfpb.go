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

	"github.com/Sirupsen/logrus"

	"github.com/pilotariak/paleta/leagues"
)

const (
	uri = "http://lidfpb.euskalpilota.fr/resultats.php"
)

var (
	challenges = map[string]string{
		"20160401": "Championnat Hiver 2016",
		"20160402": "Championnat  2015-2016",
		"20140402": "Championnat Hiver 2014",
	}

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

func (l *lidfpbLeague) Details() map[string]string {
	return map[string]string{
		"Name":    l.Name,
		"Website": l.Website,
	}
}

func (l *lidfpbLeague) Challenges() map[string]string {
	return challenges
}

func (l *lidfpbLeague) Levels() map[string]string {
	return levels
}

func (l *lidfpbLeague) Disciplines() map[string]string {
	return disciplines
}

func (l *lidfpbLeague) Display(challengeID string, disciplineID string, levelID string) error {
	logrus.Debugf("[lidfpb] Search results for %s %s %s", challengeID, disciplineID, levelID)
	return leagues.Display(uri, challengeID, disciplineID, levelID)
}
