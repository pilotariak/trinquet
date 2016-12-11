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

package storage

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/pilotariak/trinquet/config"
)

// Backend represents a storage backend
type Backend interface {

	// Name identify the key manager
	Name() string

	// Create intialize the storage backend bucket
	Create() error

	// Destroy remove the storage backend bucket
	Destroy() error

	// Put a value at the specified key
	Put(key []byte, value []byte) error

	// Get a value given its key
	Get(key []byte) ([]byte, error)

	// Delete a value given its key
	Delete(key []byte) error

	// List values
	List() ([]string, error)
}

type BackendFunc func(conf *config.Configuration) (Backend, error)

var registeredBackends = map[string](BackendFunc){}

func RegisterBackend(name string, f BackendFunc) {
	registeredBackends[name] = f
}

func New(conf *config.Configuration) (Backend, error) {
	f, ok := registeredBackends[conf.Backend]
	if !ok {
		return nil, fmt.Errorf("Unsupported backend: %s", conf.Backend)
	}
	return f(conf)
}

// GetBackends returns a list of registered storage backends
func GetBackends() []string {
	backends := make([]string, 0, len(registeredBackends))
	for name := range registeredBackends {
		backends = append(backends, name)
	}
	sort.Strings(backends)
	return backends
}

func StoreLeague(backend Backend, league *League) error {
	data, err := json.Marshal(league)
	if err != nil {
		return err
	}
	return backend.Put([]byte(league.Name), data)

}
