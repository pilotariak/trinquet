// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package api

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"github.com/pilotariak/trinquet/version"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Infof("Retrieve version")
	response := VersionResponse{
		Version: version.Version,
	}
	json.NewEncoder(w).Encode(response)
	return
}
