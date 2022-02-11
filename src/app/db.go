package app

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/timshannon/bolthold"
)

type DB interface{}

type db struct {
	ConfigDB *bolthold.Store
}

type Item struct {
	Name    string
	Created time.Time
}

func NewDB(profile string) DB {
	configPath := filepath.Join(profile, "/config/config.json")
	confDB, _ := bolthold.Open(configPath, 0666, nil)

	_ = confDB.Insert("key", &Item{
		Name:    "Test Name",
		Created: time.Now(),
	})

	var result Item
	// _ = confDB.FindOne(&result, nil)
	_ = confDB.Get("key", &result)
	fmt.Println(result)

	return &db{}
}
