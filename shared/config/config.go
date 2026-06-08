// Config parser for microservices settings
// Place the microservice config at /app/config.toml using volumes
package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config structure

type ServiceConfig struct {
	Database ServiceDatabaseConfig `toml:"database"`
}

type ServiceDatabaseConfig struct {
	User string `toml:"user"`
	Password string `toml:"password"`
	Host string `toml:"host"`
	Port uint16 `toml:"port"`
	Name string `toml:"name"`
}

// Parser

const ConfigFilePath string = "/app/config.toml"

func ParseConfig(configContent []byte) (ServiceConfig, error) {
	var cfg ServiceConfig

	if err := toml.Unmarshal(configContent, &cfg); err != nil {
		return cfg, fmt.Errorf("config file parsing error: %w", err)
	}

	return cfg, nil
}

func ParseConfigFile() (ServiceConfig, error) {
	fileData, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return ServiceConfig{}, fmt.Errorf("couldn't open the config file (%s): %w", ConfigFilePath, err)
	}

	return ParseConfig(fileData)
}
