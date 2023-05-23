package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nkolosov/whip-round/internal/app"
	"github.com/nkolosov/whip-round/internal/config"
	"github.com/nkolosov/whip-round/internal/health"
	"github.com/spf13/viper"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.New()
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

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error while loading .env file: %w", err)
	}

	viper.AutomaticEnv()

	return nil
}
