package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port            int `mapstructure:"port"`
		HealthCheckPort int `mapstructure:"health_check_port"`
	} `mapstructure:"server"`
	TokenSymmetricKey string `mapstructure:"token_symmetric_key"`
	HashSalt          string `mapstructure:"hash_salt"`
}

// New loads config from environment variables and viper config file. Returns error if config is not valid.
func New() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("server", cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("token", cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("hash", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
