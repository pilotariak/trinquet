// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const apiVersion = "1.0.0"

type model struct {
	Swagger  string `json:"swagger"`
	BasePath string `json:"basePath"`
	Info     struct {
		Title       string `json:"title"`
		Version     string `json:"version"`
		Description string `json:"description"`
	} `json:"info"`
	Schemes     []string               `json:"schemes"`
	Consumes    []string               `json:"consumes"`
	Produces    []string               `json:"produces"`
	Paths       map[string]interface{} `json:"paths"`
	Definitions map[string]interface{} `json:"definitions"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: go run main.go inputPath")
	}
	swagger := model{
		Swagger:     "2.0",
		Consumes:    []string{"application/json"},
		Produces:    []string{"application/json"},
		Paths:       make(map[string]interface{}),
		Definitions: make(map[string]interface{}),
	}
	swagger.Info.Title = "Trinquet REST API"
	swagger.Info.Version = apiVersion
	swagger.Info.Description = `
For more information about the usage of the Trinquet REST API, see
[https://github.com/pilotariak/trinquet](https://github.com/pilotariak/trinquet).
`
	fileInfos, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.Name() == "api.swagger.json" {
			continue
		}
		if !strings.HasSuffix(fileInfo.Name(), ".swagger.json") {
			continue
		}

		b, err := ioutil.ReadFile(path.Join(os.Args[1], fileInfo.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// replace "title" by "description" for fields
		b = []byte(strings.Replace(string(b), `"title"`, `"description"`, -1))

		var m model
		err = json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal(err)
		}

		for k, v := range m.Paths {
			swagger.Paths[k] = v
		}
		for k, v := range m.Definitions {
			swagger.Definitions[k] = v
		}
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(swagger)
	if err != nil {
		log.Fatal(err)
	}
}