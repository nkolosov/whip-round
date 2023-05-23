package db

import (
	"errors"
	"fmt"
	"net/url"
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
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		"postgres", url.QueryEscape(cfg.User), url.QueryEscape(cfg.Password), cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}
