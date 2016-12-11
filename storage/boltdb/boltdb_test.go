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

package boltdb

import (
	// "fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pilotariak/trinquet/config"
)

// tempfile returns a temporary file path.
func tempfile() string {
	f, _ := ioutil.TempFile("", "boltdb-")
	f.Close()
	os.Remove(f.Name())
	return f.Name()
}

// Ensure that gets a non-existent key returns nil.
func TestBoltDB_Get_NonExistent(t *testing.T) {
	db, err := NewBoltDB(&config.Configuration{
		Backend: "boltdb",
		BoltDB: &config.BoltDBConfiguration{
			Bucket: "trinquet-ut",
			File:   tempfile(),
		},
	})
	db.Create()
	if err != nil {
		t.Fatalf("Can't create BoltDB test database.")
	}
	value, err := db.Get([]byte("foo"))
	if err != nil {
		t.Fatalf("Can't retrieve BoltDB key.")
	}
	// fmt.Println("Value: ", string(value))
	if value != nil {
		t.Fatalf("Error retrieve invalid key.")
	}
}

// Ensure that that gets an existent key returns value.
func TestBoltDB_Get_Existent(t *testing.T) {
	db, err := NewBoltDB(&config.Configuration{
		Backend: "boltdb",
		BoltDB: &config.BoltDBConfiguration{
			Bucket: "trinquet-ut",
			File:   tempfile(),
		},
	})
	db.Create()
	if err != nil {
		t.Fatalf("Can't create BoltDB test database.")
	}
	db.Put([]byte("foo"), []byte("bar"))
	value, err := db.Get([]byte("foo"))
	if err != nil {
		t.Fatalf("Can't retrieve BoltDB key.")
	}
	// fmt.Println("Value: ", string(value))
	if string(value) != "bar" {
		t.Fatalf("Error retrieve invalid value.")
	}
}
