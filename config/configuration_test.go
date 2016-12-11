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

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "configuration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(templateFile.Name())
	data := []byte(`# Trinquet configuration file

# Storage backend
backend = "boltdb"

[api]
grpcPort = 8080
restPort = 9090

[boltdb]
file = "/tmp/ut.db"
bucket = "trinquet"`)
	err = ioutil.WriteFile(templateFile.Name(), data, 0700)
	if err != nil {
		t.Fatal(err)
	}
	configuration, err := LoadFileConfig(templateFile.Name())
	if err != nil {
		t.Fatalf("Error with configuration: %v", err)
	}
	fmt.Printf("Configuration : %#v\n", configuration)
	if configuration.Backend != "boltdb" {
		t.Fatalf("Configuration backend failed")
	}

	// Storage
	if configuration.BoltDB.Bucket != "trinquet" ||
		configuration.BoltDB.File != "/tmp/ut.db" {
		t.Fatalf("Configuration BoldDB failed")
	}

	// API
	if configuration.API.GrpcPort != 80 ||
		configuration.API.RestPort != 9090 {
		t.Fatalf("Configuration API failed")
	}

}
