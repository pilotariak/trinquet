// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package auth

import (
	// "fmt"
	"testing"

	"github.com/pilotariak/trinquet/config"
)

func Test_InvalidAuthenticationSystem(t *testing.T) {
	_, err := New(&config.Configuration{})
	if err == nil {
		t.Fatalf("Auth error: %s", err)
	}
}

func Test_UnknownAuthenticationSystem(t *testing.T) {
	_, err := New(&config.Configuration{
		Tracing: &config.TracingConfiguration{
			Name: "OpenBar",
		},
	})
	if err == nil {
		t.Fatal("Configure auth failed.")
	}
}
