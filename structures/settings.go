package structures

import (
	"encoding/json"
	"os"
)

// Basic settings structure
// Not exported, if you need a settings instance use the Settings function below.
type Settings struct {
	Host       string `host`
	Port       int    `port`
	Dblocation string `dblocation`
}

// Populates and returns an instance of the settings, reading from
func (s Settings) InitSettings() Settings {
	bts, err := os.ReadFile("settings.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bts, &s)
	if err != nil {
		panic(err)
	}
	return s
}
