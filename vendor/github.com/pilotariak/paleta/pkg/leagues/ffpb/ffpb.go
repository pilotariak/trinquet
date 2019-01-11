// Copyright (C) 2016-2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ffpb

import (
	"github.com/rs/zerolog/log"

	"github.com/pilotariak/paleta/pkg/leagues"
)

const (
	leagueName = "ffpb"
	uri        = "http://ffpb.euskalpilota.fr/resultats.php"
)

var (
	challenges = map[string]string{
		"20170101": "Championnat de France 2016/2017",
		"20160101": "Championnat de France 2015/2016",
		"20150101": "Championnat de France 2014/2015",
	}

	disciplines = map[string]string{
		"1":  "Trinquet / Main Nue - Groupe A",
		"2":  "Trinquet / P.G. Pleine Masculin",
		"3":  "Trinquet / P.G. Creuse Masculin",
		"4":  "Trinquet / P.G. Pleine Feminine",
		"5":  "Trinquet / P.G. Creuse Feminine",
		"6":  "Trinquet / Xare",
		"7":  "Trinquet / Paleta Pelote de Cuir",
		"8":  "Trinquet / Pasaka",
		"9":  "Trinquet / Main Nue Tête à Tête",
		"21": "Mur à Gauche / Cesta Punta",
		"22": "Mur à Gauche / Pala Corta",
		"23": "Mur à Gauche / Chistera Joko Garbi",
		"24": "Mur à Gauche / Main nue - Groupe A",
		"25": "Fronton Mur à Gauche / Paleta Pelote de Cuir",
		"26": "Mur à Gauche / P.G. Pleine Masculin",
		"28": "Mur à Gauche / P.G. Creuse Masculin Individuel",
		"29": "Fronton Mur à  Gauche / Frontenis Garçons Par équipes",
		"58": "Trinquet / Main Nue Tête à tête GrB",
		"34": "Fronton Mur à Gauche / Paleta Pelote de Gomme Pleine Corporatif",
	}
	levels = map[string]string{
		"1":  "Nationale A",
		"2":  "Nationale B",
		"4":  "Seniors",
		"5":  "Juniors",
		"6":  "Cadets",
		"7":  "Minimes",
		"8":  "Benjamins",
		"9":  "Poussins",
		"11": "Seniors 1ère Série",
		"12": "Seniors 2ème Série",
	}
)

func init() {
	leagues.RegisterLeague(leagueName, newFFPBLeague)
}

type ffpbLeague struct {
	Website     string
	Name        string
	Address     string
	Email       string
	PhoneNumber string
	Fax         string
}

func newFFPBLeague() (leagues.League, error) {
	return &ffpbLeague{
		Name:        "Fédération Française de Pelote Basque",
		Website:     "http://www.ffpb.net/",
		Address:     "Fédération Française de Pelote Basque\nBP 816 - 60, avenue Dubrocq - 64108 BAYONNE",
		Email:       "ffpbaccueil@orange.fr",
		PhoneNumber: "05.59.59.22.34",
		Fax:         "05.59.25.49.82",
	}, nil
}

func (l *ffpbLeague) Details() map[string]string {
	return map[string]string{
		"Name":        l.Name,
		"Website":     l.Website,
		"Address":     l.Address,
		"Email":       l.Email,
		"PhoneNumber": l.PhoneNumber,
		"Fax":         l.Fax,
	}
}

func (l *ffpbLeague) Challenges() map[string]string {
	return challenges
}

func (l *ffpbLeague) Levels() map[string]string {
	return levels
}

func (l *ffpbLeague) Disciplines() map[string]string {
	return disciplines
}

func (l *ffpbLeague) Display(challengeID string, disciplineID string, levelID string) error {
	log.Debug().Str("league", leagueName).Msgf("Search results for %s %s %s", challengeID, disciplineID, levelID)
	return leagues.Display(uri, challengeID, disciplineID, levelID)
}
