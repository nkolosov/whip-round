package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"github.com/nkolosov/whip-round/internal/hash"
	"github.com/nkolosov/whip-round/internal/repository"
	"github.com/nkolosov/whip-round/internal/token"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user with email %s already exists", "%s")
)

type UserService struct {
	repo         repository.User
	hashManager  hash.HashManager
	tokenManager token.Manager
}

func NewUserService(repo repository.User, tokenManager token.Manager, hashManager hash.HashManager) *UserService {
	return &UserService{
		repo:         repo,
		hashManager:  hashManager,
		tokenManager: tokenManager,
	}
}

func (s *UserService) CreateUser(ctx context.Context, userDTO *domain.UserDTO) (*domain.User, error) {
	if userDTO.Validate() != nil {
		return nil, errors.New("user validation failed")
	}

	if _, err := s.repo.GetUserByEmail(ctx, userDTO.Email); err == nil {
		return nil, fmt.Errorf(ErrUserAlreadyExists.Error(), userDTO.Email)
	}

	user, err := domain.ConvertUserDTOToUser(userDTO)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New()

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
