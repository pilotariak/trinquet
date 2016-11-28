# Paleta

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/pilotariak%2Fpaleta.svg)](https://badge.fury.io/gh/pilotariak%2Fpaleta)

Master :
* [![Circle CI](https://circleci.com/gh/pilotariak/paleta/tree/master.svg?style=svg)](https://circleci.com/gh/pilotariak/paleta/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/pilotariak/paleta/tree/develop.svg?style=svg)](https://circleci.com/gh/pilotariak/paleta/tree/develop)

This tool is a simple CLI to display informations about Pelota competitions for the terminal.
Supported leagues are :

* [x] [Côte d'Argent](http://www.lcapb.net/)
* [x] [Pays Basque](http://www.comite-pelote-basque.eus/fr/)
* [ ] [Landes](http://www.llpb.fr/)
* [ ] [Béarn](http://liguebearnpelote.fr/)
* [ ] [Midi-Pyrénées](http://www.lmppb.fr/)
* [x] [Ile de France](http://www.lidfpb.fr/)

* [x] [Fédération française](http://www.ffpb.net/)


![Screenshot](paleta.png)


## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/paleta-0.8.0_netbsd_arm) ]


## Usage

* Show supported leagues :

        $ ./paleta -leagues
        Leagues:
        - ctpb
        - lbpb
        - lcapb
        - lidfpb

* Show available levels for a league :

        $ ./paleta -league lcapb -levels
        Levels:
        - [51] Senoir Individuel
        - [1] 1ère Série
        - [2] 2ème Série
        - [3] 3ème Série
        - [4] Seniors
        - [6] Cadets
        - [9] Poussins
        - [7] Minimes
        - [8] Benjamins

* Show available disciplines for a league:

        $ ./paleta -league lcapb -disciplines
        Disciplines:
        - [2] Trinquet / P.G. Pleine Masculin
        - [4] Trinquet / P.G. Pleine Feminine
        - [26] Mur à Gauche / P.G. Pleine Masculin
        - [126] Mur A gauche / P.G. Pleine Masculin Barrages
        - [501] Place Libre / P.G Pleine Feminine
        - [3] Trinquet / P.G. Creuse Masculin
        - [5] Trinquet / P.G. Creuse Feminine
        - [13] Place Libre / Grand Chistera
        - [16] Place Libre / P.G. Pleine Masculin
        - [27] Mur à Gauche / P.G. Pleine Feminine
        - [28] Mur à Gauche / P.G. Creuse Masculin Individuel

* Display result for a competiion:

        $ ./paleta -league lcapb -level 1 -discipline 2
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        |     DATE      |               CLUB 1               |               CLUB 2               | SCORE | COMMENTAIRE |
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        | Poules 1      | Club 1                             | Club 2                             | Score | Commentaire |
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        | 18/09/2016    | PILOTARI IRRATZABAL CLUB           | PILOTARI IRRATZABAL CLUB           | 39/40 |             |
        |               | (040802 - 0120) DAINCIART Jon (E)  | (015900 - 0571) IRIART Laurent (E) |       |             |
        |               | (039742) CAZAUBON Emerik           | (016667) BORDACHAR Serge           |       |             |
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        | 18/09/2016    | C.A. BEGLAIS                       | A.P. AVIATION CIVILE & METEO. BX   | 40/31 |             |
        |               | (073842) GARCIA Antoine (S)        | (053104) BEDECARRAX Patrice        |       |             |
        |               | (073197) SENDER Timothe (S)        | (055181) DE Bouyn Godefroy         |       |             |
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        ...
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        | Finale        | Club 1                             | Club 2                             | Score | Commentaire |
        +---------------+------------------------------------+------------------------------------+-------+-------------+
        | 27/11/2016    | C.A. BEGLAIS                       | PILOTARI IRRATZABAL CLUB           | 40/36 |             |
        |               | (073842) GARCIA Antoine (S)        | (015900 - 0571) IRIART Laurent (E) |       |             |
        |               | (073197) SENDER Timothe (S)        | (016667) BORDACHAR Serge           |       |             |
        +---------------+------------------------------------+------------------------------------+-------+-------------+



## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
