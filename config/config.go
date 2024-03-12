package config

import (
	"encoding/json"
	"log"
	"os"
)

const ConfigFilePath string = ".joven-data.json"

// Config represents the configuration settings for the application.
type Config struct {
	Token  string   `json:"token"`  // token is the authentication token for the GitLab.
	Groups []string `json:"groups"` // groups is a list of group names for the GitLAb.
}

func constructConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}
	return home + "/" + ConfigFilePath

}

func New(token string, groups []string) *Config {
	return &Config{Token: token, Groups: groups}
}

func (c *Config) Save() error {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	filePath := constructConfigPath()

	err = os.WriteFile(filePath, config, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Load() (*Config, error) {
	filePath := constructConfigPath()

	configData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}
