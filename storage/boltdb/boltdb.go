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
	"fmt"
	"os"
	"path"

	"github.com/boltdb/bolt"
	"github.com/golang/glog"

	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/storage"
)

const (
	// Bucket defines the Trinquet bucket
	// Bucket      = "trinquet"
	boltdbLabel = "boltdb"
)

func init() {
	storage.RegisterBackend(boltdbLabel, NewBoltDB)
}

// BoltDB is the Boltdb backend.
type BoltDB struct {
	*bolt.DB
	BucketName string
	Path       string
}

// NewBoltDB opens a new BoltDB connection to the specified path and bucket
func NewBoltDB(conf *config.Configuration) (storage.Backend, error) {
	//f := filepath.Join(conf.BoltDB.Directory, "trinquet.db")
	f := conf.BoltDB.File
	directory := path.Dir(f)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		glog.V(1).Infof("Create %s", directory)
		err = os.Mkdir(directory, 0755)
		if err != nil {
			return nil, err
		}
	}
	db, err := bolt.Open(f, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltDB{
		DB:         db,
		Path:       conf.BoltDB.File,
		BucketName: conf.BoltDB.Bucket,
	}, nil
}

// Name returns BoltDB label
func (db *BoltDB) Name() string {
	return boltdbLabel
}

// Create intialize the storage backend
func (db *BoltDB) Create() error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(db.BucketName))
		if err != nil {
			return fmt.Errorf("BoltDB Can't create bucket: %s", err)
		}
		return nil
	})
}

// Destroy remove the storage backend bucket
func (db *BoltDB) Destroy() error {
	directory := path.Dir(db.BucketName)
	glog.V(1).Infof("Create %s", directory)
	return nil
}

// List returns all secrets
func (db *BoltDB) List() ([]string, error) {
	glog.V(1).Infof("List entries")
	var l []string
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.ForEach(func(key, value []byte) error {
			keyData := string(key)
			valueData := string(value)
			glog.V(2).Infof("Entry: %s %s", keyData, valueData)
			l = append(l, keyData)
			return nil
		})
		return nil
	})
	return l, nil
}

// Get a value given its key
func (db *BoltDB) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	var value []byte
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.ForEach(func(k, v []byte) error {
			//log.Printf("[BoltDB] Entry : %s %s", string(k), string(v))
			if string(k) == string(key) {
				glog.V(1).Infof("Find : %s", string(v))
				value = v
				return nil
			}
			return nil
		})
		return nil
	})
	return value, nil
}

// Put a value at the specified key
func (db *BoltDB) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	return db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.Put(key, value)
		return nil
	})
}

// Delete the value at the specified key
func (db *BoltDB) Delete(key []byte) error {
	glog.V(1).Infof("Delete : %v", string(key))
	return fmt.Errorf("Not implemented")
}
