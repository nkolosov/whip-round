package db

import (
	"errors"
	"fmt"
)

var (
	ErrConfigIsNil = errors.New("config is nil")
)

type Config struct {
	Host         string
	Port         int
	User         string
	DBName       string
	SSLMode      string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

func (cfg *Config) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password)
}
