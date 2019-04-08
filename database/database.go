// Package database takes care of read / write process.
package database

import (
	"io/ioutil"
	"log"
	"os"
)

// Database describes accesses to a storage
type Database interface {
	Read() []byte
	Write([]byte) error
}

// File is a storage onto the disk
type File struct{}

// Read fetches data from storage
func (f File) Read() []byte {
	rawJSON, err := os.Open(f.path())
	if err != nil {
		log.Fatal(err)
	}
	defer rawJSON.Close()

	readJSON, err2 := ioutil.ReadAll(rawJSON)
	if err2 != nil {
		log.Fatal(err2)
	}

	return readJSON
}

// Write effectively writes data into storage
func (f File) Write(data []byte) error {
	return ioutil.WriteFile(f.path(), data, 0)
}

func (f File) path() string {
	p := "/src/github.com/bruno-chavez/restedancestor/database/database.json"
	return os.Getenv("GOPATH") + p
}
