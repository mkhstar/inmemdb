package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DB struct that holds the data
type db struct {
	mu           *sync.Mutex
	data         map[string]interface{}
	keysToExpire map[string]time.Time
}

type dset map[string]struct{}

type dlist []string

var inMemDB = db{
	mu:           &sync.Mutex{},
	data:         make(map[string]interface{}),
	keysToExpire: make(map[string]time.Time),
}

func init() {
	go keyCleaner()
	Backup()
}

// Persist persists inMemDB to disk
func Persist() {
	dataToPersist := map[string]interface{}{
		"data":         inMemDB.data,
		"keysToExpire": inMemDB.keysToExpire,
	}

	jsonData, _ := json.Marshal(dataToPersist)

	rootPath, _ := os.Executable()

	ioutil.WriteFile(filepath.Join(rootPath, "..", "data.json"), []byte(jsonData), 0666)
}

// Backup recovers inMemDB from disk
func Backup() {
	rootPath, err := os.Executable()

	if err != nil {
		log.Println("Failed to backup due to error", err)
	} else {
		var store Store
		data, err := ioutil.ReadFile(filepath.Join(rootPath, "..", "data.json"))

		if err == nil {
			json.Unmarshal(data, &store)
			inMemDB.data = store.Data
			inMemDB.keysToExpire = store.KeysToExpire
		}

	}

}
