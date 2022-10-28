package cmd

import (
	"encoding/json"
	"os"
)

// Configuration is a struct designed to hold the applications variable configuration settings
type Configuration struct {
	Port           string
	APIHost        string
	SessionManager string
	RedisHost      string
	RedisPassword  string
	ENV            string
}

// ConfigurationSettings is a function that reads a json configuration file and outputs a Configuration struct
func ConfigurationSettings(env string) (*Configuration, error) {
	confFile := "conf.json"
	if env == "test" {
		confFile = "test_conf.json"
	}
	file, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	configurationSettings := Configuration{}
	err = decoder.Decode(&configurationSettings)
	if err != nil {
		return nil, err
	}
	return &configurationSettings, nil
}

// InitializeEnvironment sets environmental variables
func (c *Configuration) InitializeEnvironment() {
	os.Setenv("PORT", c.Port)
	os.Setenv("API_HOST", c.APIHost)
	os.Setenv("SESSION_MANAGER", c.SessionManager)
	os.Setenv("REDIS_HOST", c.RedisHost)
	os.Setenv("REDIS_PASSWORD", c.RedisPassword)
	os.Setenv("ENV", c.ENV)
}

func InitializeEnvironment() {
	os.Setenv("PORT", "8080")
	os.Setenv("API_HOST", "http://localhost:8081")
	os.Setenv("SESSION_MANAGER", "Default")
	os.Setenv("REDIS_HOST", "localhost:55000")
	os.Setenv("REDIS_PASSWORD", "redispw")
	os.Setenv("ENV", "dev")
}
