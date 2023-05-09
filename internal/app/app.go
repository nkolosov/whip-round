package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nkolosov/whip-round/internal/config"
	"github.com/nkolosov/whip-round/internal/db/memory"
	"github.com/nkolosov/whip-round/internal/handlers"
	"github.com/nkolosov/whip-round/internal/hash"
	"github.com/nkolosov/whip-round/internal/repository"
	"github.com/nkolosov/whip-round/internal/server"
	"github.com/nkolosov/whip-round/internal/services"
	"github.com/nkolosov/whip-round/internal/token"
	"net/http"
	"time"
)

func App(cfg *config.Config) (*server.Server, error) {
	store, err := memory.NewMemoryDB()
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

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
