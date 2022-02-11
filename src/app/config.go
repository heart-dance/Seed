package app

import (
	"github.com/timshannon/bolthold"
)

type Config struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	Profile       string `json:"profile"`
	ProfileData   string `json:"profile_data"`
	ProfileCache  string `json:"profile_cache"`
	ProfileConfig string `json:"profile_config"`
}

func GetConfigDB() {
	bolthold.Open("", 0666, nil)
}
