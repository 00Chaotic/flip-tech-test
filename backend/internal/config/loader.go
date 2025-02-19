package config

import (
	"github.com/kelseyhightower/envconfig"
)

// LoadConfig parses the Config object from environment variables.
func LoadConfig() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
