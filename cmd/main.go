package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nkolosov/whip-round/internal/app"
	"github.com/nkolosov/whip-round/internal/config"
	"github.com/nkolosov/whip-round/internal/health"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := app.App(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(srv.Run())
	}()

	go health.StartHealthCheckServer(cfg.Server.HealthCheckPort)

	log.Println("Starting server on port ", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Server shutting down...")

	log.Fatal(srv.Shutdown(context.Background()))
}

func initConfig() (*config.Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error while loading .env file: %w", err)
	}

	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("error while reading config: %w", err)
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return nil, fmt.Errorf("error while reading port from env: %w", err)
	}
	cfg.Server.Port = port

	healthCheckPort, err := strconv.Atoi(os.Getenv("HEALTH_CHECK_PORT"))
	if err != nil {
		return nil, fmt.Errorf("error while reading health check port from env: %w", err)
	}
	cfg.Server.HealthCheckPort = healthCheckPort

	cfg.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
	cfg.HashSalt = os.Getenv("HASH_SALT")

	cfg.DB.Host = os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.DB.Port = dbPort
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.DBName = os.Getenv("DB_NAME")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.SSLMode = os.Getenv("DB_SSL_MODE")
	cfg.DB.MaxConn = os.Getenv("DB_MAX_OPEN_CONN")

	fmt.Printf("%+v\n", cfg)

	return cfg, nil
}
