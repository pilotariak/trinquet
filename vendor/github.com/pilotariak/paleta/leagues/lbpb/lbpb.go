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

package lbpb

import (
	// "fmt"

	"github.com/Sirupsen/logrus"

	"github.com/pilotariak/paleta/leagues"
)

const (
	uri = "http://lbpb.euskalpilota.fr/resultats.php"
)

var (
	challenges = map[string]string{
		"20170301": "CHAMPIONNAT HIVER 2016 - 2017",
		"20160301": "Championnat HIVER 2015 - 2016",
		"20140302": "Championnat HIVER 2014 - 2015",
	}

	disciplines = map[string]string{
		"2":  "Trinquet / P.G. Pleine Masculin",
		"3":  "Trinquet / P.G. Creuse Masculin",
		"4":  "Trinquet / P.G. Pleine Feminine",
		"5":  "Trinquet / P.G. Creuse Feminine",
		"6":  "Trinquet / Xare",
		"7":  "Trinquet / Paleta Cuir",
		"12": "Place Libre / Chistera Joko Garbi",
		"22": "Mur à Gauche / Pala Corta",
		"23": "Mur à Gauche / Chistera Joko Garbi",
		"26": "Mur à Gauche / P.G. Pleine Masculin",
		"28": "Mur à Gauche / P.G. Creuse Masculin Individuel",
		"29": "Mur à Gauche / Frontenis Féminin",
		"30": "Mur à Gauche / Frontenis Masculin",
		"41": "Mur à Gauche / Paleta Cuir Jeunes",
		"31": "Mur à Gauche / P.G Creuse Féminine",
		"32": "Mur à Gauche / P.G Creuse Masculin par éq.",
	}

	levels = map[string]string{
		"1":  "1ère Série",
		"2":  "2ème Série",
		"3":  "3ème Série",
		"4":  "Seniors",
		"6":  "Cadets",
		"7":  "Minimes",
		"8":  "Benjamins",
		"9":  "Poussins",
		"10": "Minimes 2",
		"11": "Benjamins 2",
		"12": "Poussins 2",
	}
)

func init() {
	leagues.RegisterLeague("lbpb", newLBPBLeague)
}

type lbpbLeague struct {
	Website     string
	Name        string
	Address     string
	Email       string
	PhoneNumber string
	Fax         string
}

func newLBPBLeague() (leagues.League, error) {
	return &lbpbLeague{
		Name:        "Ligue du Béarn de Pelote",
		Website:     "http://liguebearnpelote.fr/lbp/",
		Address:     "Centre Nelson Paillou\n12 rue du Professeur Garrigou-Lagrange\n64000 PAU",
		Email:       "liguebearnpelote@wanadoo.fr",
		PhoneNumber: "05 59 14 19 98 ",
		Fax:         "05 59 14 19 99",
	}, nil
}

func (l *lbpbLeague) Details() map[string]string {
	return map[string]string{
		"Name":        l.Name,
		"Website":     l.Website,
		"Address":     l.Address,
		"Email":       l.Email,
		"PhoneNumber": l.PhoneNumber,
		"Fax":         l.PhoneNumber,
	}
}

func (l *lbpbLeague) Challenges() map[string]string {
	return challenges
}

func (l *lbpbLeague) Levels() map[string]string {
	return levels
}

func (l *lbpbLeague) Disciplines() map[string]string {
	return disciplines
}

func (l *lbpbLeague) Display(challengeID string, disciplineID string, levelID string) error {

	logrus.Debugf("[lbpb] Search results for %s %s %s", challengeID, disciplineID, levelID)
	return leagues.Display(uri, challengeID, disciplineID, levelID)
}
