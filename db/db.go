// Copyright © 2018 zergwangj <zergwangj@163.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"github.com/boltdb/bolt"
	"fmt"
)

const dbFile = "zergpass.db"
const passwordsBucket = "passwords"

type DB struct {
	Db			*bolt.DB
}

func NewDB() *DB {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b == nil {
			fmt.Println("No existing passwords found. Creating a new one...")
			_, err := tx.CreateBucket([]byte(passwordsBucket))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil
	}

	d := &DB {
		Db:  		db,
	}
	return d
}

func (d *DB) Close() {
	d.Db.Close()
}

func (d *DB) Add(entry *Entry) error {
	err := d.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			return nil
		}
		return fmt.Errorf("Bucket passwords is not found")
	})
	if err != nil {
		return err
	}

	err = d.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			data, err := entry.Serialize()
			if err != nil {
				return err
			}
			err = b.Put([]byte(entry.Title), data)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Delete(title string) error {
	err := d.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			data := b.Get([]byte(title))
			if data == nil {
				return fmt.Errorf("Entry %s not found", title)
			}
			return nil
		}
		return fmt.Errorf("Bucket passwords is not found")
	})
	if err != nil {
		return err
	}

	err = d.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			err = b.Delete([]byte(title))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Set(entry *Entry) error {
	err := d.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			data := b.Get([]byte(entry.Title))
			if data == nil {
				return fmt.Errorf("Entry %s not found", entry.Title)
			}
			return nil
		}
		return fmt.Errorf("Bucket passwords is not found")
	})
	if err != nil {
		return err
	}

	err = d.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			data, err := entry.Serialize()
			if err != nil {
				return err
			}
			err = b.Put([]byte(entry.Title), data)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Get(title string) (*Entry, error) {
	var entry *Entry
	var err error

	err = d.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			data := b.Get([]byte(title))
			if data == nil {
				return fmt.Errorf("Entry %s not found", title)
			}
			entry, err = DeserializeEntry(data)
			if err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("Bucket passwords is not found")
	})
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (d *DB) List() ([]*Entry, error) {
	entries := make([]*Entry, 0, 100)

	err := d.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(passwordsBucket))
		if b != nil {
			b.ForEach(func(k, v []byte) error {
				entry, err := DeserializeEntry(v)
				if err != nil {
					return err
				}
				entries = append(entries, entry)
				return nil
			})
			return nil
		}
		return fmt.Errorf("Bucket passwords is not found")
	})
	if err != nil {
		return nil, err
	}

	return entries, nil
}
