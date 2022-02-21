package model

type Config struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	Profile       string `json:"profile"`
	ProfileData   string `json:"profile_data"`
	ProfileCache  string `json:"profile_cache"`
	ProfileConfig string `json:"profile_config"`
}
