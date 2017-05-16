package api

import (
	//"fmt"
	"net/http"

	"github.com/golang/glog"

	"github.com/pilotariak/trinquet/pkg/static"
	"github.com/pilotariak/trinquet/pkg/webdoc"
)

// ServeStaticFile expose static files
func ServeStaticFile(mux *http.ServeMux) {
	glog.V(1).Infof("Create the Static file handler")

	mux.HandleFunc("/changelog", func(w http.ResponseWriter, req *http.Request) {
		if data, err := static.ChangelogMdBytes(); err == nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write(data)
		}
	})
	mux.HandleFunc("/doc", func(w http.ResponseWriter, req *http.Request) {
		if data, err := webdoc.DocIndexHtmlBytes(); err == nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(data)
		}
	})
}
