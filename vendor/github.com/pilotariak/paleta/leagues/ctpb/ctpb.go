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
	uri = "http://ctpb.euskalpilota.fr/resultats.php"
)

var (
	challenges = map[string]string{
		"20170501": "Championnat d'Hiver 2016-2017",
		"20160104": "Championnat d'Hiver 2015-2016",
		"20150104": "Championnat d'Hiver 2014-2015",
		"20140104": "Championnat d'Hiver 2013-2014",
		"20130104": "Championnat d'Hiver 2012-2013",
	}

	disciplines = map[string]string{
		"1":   "Trinquet / Main Nue - Groupe A",
		"2":   "Trinquet / P.G. Pleine Masculin",
		"3":   "Trinquet / P.G. Creuse Masculin",
		"4":   "Trinquet / P.G. Pleine Feminine",
		"5":   "Trinquet / P.G. Creuse Feminine",
		"6":   "Trinquet / Xare",
		"7":   "Trinquet / Paleta Cuir",
		"8":   "Trinquet / Pasaka",
		"9":   "Trinquet / Main Nue Tête à Tête",
		"13":  "Place Libre / Grand Chistera",
		"16":  "Place Libre / P.G. Pleine Masculin",
		"21":  "Mur à Gauche / Cesta Punta",
		"22":  "Mur à Gauche / Pala Corta",
		"23":  "Mur à Gauche / Chistera Joko Garbi",
		"24":  "Mur à Gauche / Main nue - Groupe A",
		"26":  "Mur à Gauche / P.G. Pleine Masculin",
		"27":  "Mur à Gauche / P.G. Pleine Feminine",
		"28":  "Mur à Gauche / P.G. Creuse Masculin Individuel",
		"49":  "Trinquet / PASAKA - Coupe Lemoine",
		"57":  "Trinquet / Main Nue - Groupe B",
		"58":  "Trinquet / Main Nue Tête à tête GrB",
		"60":  "Mur à Gauche / Main nue - Groupe B",
		"105": "Mur à Gauche / Cesta Punta - Groupe B",
		"115": "Mur à Gauche / Joko Garbi - Groupe B",
		"116": "Trinquet / PG Pleine Féminine - Groupe B",
		"126": "Mur A gauche / P.G. Pleine Masculin Barrages",
		"501": "Place Libre / P.G Pleine Feminine",
	}
	levels = map[string]string{
		"1": "1ère Série - 1.Maila",
		"2": "2ème Série - 2.Maila",
		"3": "3ème Série - 3.Maila",
		"5": "Juniors - Artekoak",
		"6": "Cadets - Gazteak",
		"7": "Minimes - Gaztetxoak",
		"8": "Benjamins - Kimuak",
		"9": "Poussins - Umetxoak",
	}
)

func init() {
	leagues.RegisterLeague("ctpb", newCTPBLeague)
}

type ctpbLeague struct {
	Website     string
	Name        string
	Address     string
	Email       string
	PhoneNumber string
	Fax         string
}

func newCTPBLeague() (leagues.League, error) {
	return &ctpbLeague{
		Website:     "http://www.comite-pelote-basque.eus/fr/",
		Name:        "Comité Territorial Pays Basque de Pelote Basque",
		Address:     "7, place du jeu de paume\n64240 HASPARREN",
		Email:       "info@comite-pelote-basque.eus",
		Fax:         "05-59-29-49-61",
		PhoneNumber: "05-59-29-59-40",
	}, nil
}

func (l *ctpbLeague) Details() map[string]string {
	return map[string]string{
		"Name":        l.Name,
		"Website":     l.Website,
		"Address":     l.Address,
		"Email":       l.Email,
		"PhoneNumber": l.PhoneNumber,
		"Fax":         l.Fax,
	}
}

func (l *ctpbLeague) Challenges() map[string]string {
	return challenges
}

func (l *ctpbLeague) Levels() map[string]string {
	return levels
}

func (l *ctpbLeague) Disciplines() map[string]string {
	return disciplines
}

func (l *ctpbLeague) Display(challengeID string, disciplineID string, levelID string) error {
	logrus.Debugf("[ctpb] Search results for %s %s %s", challengeID, disciplineID, levelID)
	return leagues.Display(uri, challengeID, disciplineID, levelID)
}
