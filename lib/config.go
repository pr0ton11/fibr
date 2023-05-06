package lib

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Configuration
type Config struct {
	// URL of the server
	server *string `yaml:"server"`
	// Heartbeat configuration
	heartbeat *HeartBeatConfiguration `yaml:"heartbeat"`
	// Models contained in this configuration
	models *[]LanguageModel `yaml:"languageModels"`
}

// Constructor for new configuration
func NewConfig() *Config {
	return &Config{}
}

// Loads configuration from YAML file
func (config *Config) load(path string) error {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read configuration file %s:\n %w", path, err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("could not parse configuration file %s:\n %w", path, err)
	}
	return nil
}
