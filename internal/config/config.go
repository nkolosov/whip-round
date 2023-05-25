package config

import "github.com/spf13/viper"

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
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DBName         string `mapstructure:"db_name"`
	SSLMode        string `mapstructure:"ssl_mode"`
	MaxOpenConns   int    `mapstructure:"max_open_conns"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns"`
	ConnectTimeout int    `mapstructure:"connect_timeout"`
}

func New() (*Config, error) {
	cfg := &Config{}

	cfg.Server.Port = viper.GetInt("APP_PORT")
	cfg.Server.HealthCheckPort = viper.GetInt("HEALTH_CHECK_PORT")
	cfg.TokenSymmetricKey = viper.GetString("TOKEN_SYMMETRIC_KEY")
	cfg.HashSalt = viper.GetString("HASH_SALT")
	cfg.DB.Host = viper.GetString("DB_HOST")
	cfg.DB.Port = viper.GetInt("DB_PORT")
	cfg.DB.User = viper.GetString("DB_USER")
	cfg.DB.Password = viper.GetString("DB_PASSWORD")
	cfg.DB.DBName = viper.GetString("DB_NAME")
	cfg.DB.SSLMode = viper.GetString("DB_SSL_MODE")
	cfg.DB.MaxOpenConns = viper.GetInt("DB_MAX_OPEN_CONNS")
	cfg.DB.MaxIdleConns = viper.GetInt("DB_MAX_IDLE_CONNS")
	cfg.DB.ConnectTimeout = viper.GetInt("DB_CONNECT_TIMEOUT")

	return cfg, nil
}
