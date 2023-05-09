package services

import (
	"errors"
	"github.com/nkolosov/whip-round/internal/hash"
	"github.com/nkolosov/whip-round/internal/repository"
	"github.com/nkolosov/whip-round/internal/token"
)

////go:generate mockgen -destination=mocks/mocks.go -package=mocks github.com/nkolosov/whip-round/internal/services ServiceI,UserServiceI,SessionServiceI

type Service struct {
	AuthService *AuthService
	UserService *UserService
}

func NewService(repository *repository.Repository, tokenManager token.Manager, hashManager hash.HashManager) (*Service, error) {
	if repository == nil {
		return nil, errors.New("repository is nil")
	}

	if tokenManager == nil {
		return nil, errors.New("tokenManager is nil")
	}

	return &Service{
		AuthService: NewAuthService(repository, tokenManager, hashManager),
		UserService: NewUserService(repository, tokenManager, hashManager),
	}, nil
}
