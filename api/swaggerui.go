// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package api

import (
	"io"
	"strings"
	//"fmt"
	"mime"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"

	"github.com/pilotariak/trinquet/pb"
	"github.com/pilotariak/trinquet/pkg/ui/swagger"
)

const (
	prefix = "/swagger-ui/"
)

// ServeSwagger expose files in third_party/swagger-ui/ on <host>/swagger-ui
func ServeSwagger(mux *http.ServeMux) {
	glog.V(1).Infof("Create the SwaggerUI handler")
	mime.AddExtensionType(".svg", "image/svg+xml")

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:     swagger.Asset,
		AssetDir:  swagger.AssetDir,
		AssetInfo: swagger.AssetInfo,
		Prefix:    "third_party/swagger-ui",
	})
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
