package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port            int `mapstructure:"port"`
		HealthCheckPort int `mapstructure:"health_check_port"`
	} `mapstructure:"server"`
	DB                PostgresConfig
	TokenSymmetricKey string `mapstructure:"token_symmetric_key"`
	HashSalt          string `mapstructure:"hash_salt"`
}

type PostgresConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	SSLMode      string `mapstructure:"ssl_mode"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

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

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	return cfg, nil
}
