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
	"bytes"
	"encoding/gob"
	"reflect"
)

type Entry struct {
	Title 			string
	Url 			string
	Username 		string
	Password 		string
	Notes 			string
}

func NewEntry() *Entry {
	return &Entry{}
}

func FieldStrings() []string {
	strings := make([]string, 0)
	entryType := reflect.TypeOf(Entry{})
	for i := 0; i < entryType.NumField(); i++ {
		strings = append(strings, entryType.Field(i).Name)
	}
	return strings
}

func (e *Entry) ValueStrings() []string {
	strings := make([]string, 0)
	strings = append(strings, e.Title)
	strings = append(strings, e.Url)
	strings = append(strings, e.Username)
	strings = append(strings, e.Password)
	strings = append(strings, e.Notes)
	return strings
}

func (e *Entry) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(e)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func DeserializeEntry(d []byte) (*Entry, error) {
	var entry Entry

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}