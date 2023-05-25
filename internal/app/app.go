package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"

	"github.com/nkolosov/whip-round/internal/config"
	"github.com/nkolosov/whip-round/internal/db"
	"github.com/nkolosov/whip-round/internal/db/postgres"
	"github.com/nkolosov/whip-round/internal/handlers"
	"github.com/nkolosov/whip-round/internal/hash"
	"github.com/nkolosov/whip-round/internal/repository"
	"github.com/nkolosov/whip-round/internal/server"
	"github.com/nkolosov/whip-round/internal/services"
	"github.com/nkolosov/whip-round/internal/token"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func App(cfg *config.Config) (*server.Server, error) {
	if err := runMigration(cfg); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store, err := postgres.NewPool(ctx, &db.Config{
		Host:           cfg.DB.Host,
		Port:           cfg.DB.Port,
		User:           cfg.DB.User,
		DBName:         cfg.DB.DBName,
		Password:       cfg.DB.Password,
		SSLMode:        cfg.DB.SSLMode,
		MaxIdleConns:   cfg.DB.MaxIdleConns,
		MaxOpenConns:   cfg.DB.MaxOpenConns,
		ConnectTimeout: cfg.DB.ConnectTimeout,
	})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer store.Close()

	repo, err := repository.NewRepository(store)
	if err != nil {
		return nil, fmt.Errorf("failed to create repo: %w", err)
	}

	tokenManager, err := token.NewPasetoManager(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	hashManager, err := hash.NewHash(cfg.HashSalt)
	if err != nil {
		return nil, fmt.Errorf("failed to create hash manager: %w", err)
	}

	service, err := services.NewService(repo, tokenManager, hashManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	r := gin.Default()
	h := handlers.NewHandler(service)

	return server.NewServer(&http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        h.Init(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}), nil
}

func runMigration(cfg *config.Config) error {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode)

	m, err := migrate.New("file://schema", dbConn)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	return nil
}
