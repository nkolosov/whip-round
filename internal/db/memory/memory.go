package memory

import (
	"context"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
)

type UserDB interface {
	CreateUser(ctx context.Context, u *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type SessionDB interface {
	CreateRefreshToken(ctx context.Context, refreshToken *domain.RefreshSession) error
	GetRefreshToken(ctx context.Context, tokenValue string) (*domain.RefreshSession, error)
	DeleteRefreshToken(ctx context.Context, tokenValue string) error
	DeleteRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) error
}

////go:generate mockgen -destination=mocks/mock.go -package=mocks github.com/nkolosov/whip-round/internal/db/memory UserDB,SessionDB,MemoryDB

type DB struct {
	userStore    map[string]*domain.User
	sessionStore []*domain.RefreshSession
}

func NewMemoryDB() (*DB, error) {
	return &DB{
		sessionStore: make([]*domain.RefreshSession, 0),
		userStore:    make(map[string]*domain.User),
	}, nil
}
