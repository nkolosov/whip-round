package services

import (
	"context"
	"time"

	"github.com/nkolosov/whip-round/internal/hash"
	"github.com/nkolosov/whip-round/internal/repository"
	"github.com/nkolosov/whip-round/internal/token"
)

type AuthService struct {
	repo         repository.User
	hashManager  hash.HashManager
	tokenManager token.Manager
}

func NewAuthService(repo repository.User, tokenManager token.Manager, hashManager hash.HashManager) *AuthService {
	return &AuthService{
		repo:         repo,
		hashManager:  hashManager,
		tokenManager: tokenManager,
	}
}

// GetAccessToken returns access token and refresh token for specific email and password.
// It returns error if any.
func (s *AuthService) GetAccessToken(ctx context.Context, email string, password string) (accessToken string, refreshToken string, err error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	// TODO: uncomment this
	//if !s.hashManager.CheckPasswordHash(password, user.Password) {
	//	return "", "", fmt.Errorf("invalid password")
	//}

	accessToken, err = s.tokenManager.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.tokenManager.CreateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
