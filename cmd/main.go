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

	return cfg, nil
}
