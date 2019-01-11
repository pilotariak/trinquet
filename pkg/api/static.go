// Copyright (C) 2016-2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"fmt"
	"html/template"
	"mime"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/rs/zerolog/log"

	"github.com/pilotariak/trinquet/pkg/changelog"
	"github.com/pilotariak/trinquet/pkg/static"
	"github.com/pilotariak/trinquet/pkg/version"
	"github.com/pilotariak/trinquet/pkg/web"
)

type homePage struct {
	Version string
}

// ServeStaticFile expose static files
func ServeStaticFile(mux *http.ServeMux) {
	log.Info().Msg("Create the Static file handler")

	mux.HandleFunc("/changelog", func(w http.ResponseWriter, req *http.Request) {
		log.Info().Str("api", "changelog").Msg("Create the Static file handler")
		if data, err := changelog.Asset("ChangeLog.md"); err == nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write(data)
		}
	})

	log.Debug().Msg("Create the media handler")
	// log.Debug().Msg("Media files : %s", static.AssetNames())
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".png", "image/png")
	mime.AddExtensionType(".svg", "image/svg+xml")
	mux.Handle("/static/", http.FileServer(&assetfs.AssetFS{
		Asset:     static.Asset,
		AssetDir:  static.AssetDir,
		AssetInfo: static.AssetInfo,
		// Prefix:    "css",
	}))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tplFile, err := web.Asset("templates/home.tpl")
		if err != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(fmt.Sprintf("Can't load home template: %s", err)))
			return
		}
		t, err := template.New("home").Parse(string(tplFile))
		if err != nil {
			log.Error().Err(err).Msg("Status can't parse template")
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(fmt.Sprintf("Can't render home: %s", err)))
			return
		}
		page := &homePage{
			Version: version.Version,
		}
		log.Info().Msg("Rendering Home page")
		t.Execute(w, page)
	})

}
