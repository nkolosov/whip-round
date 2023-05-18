package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
)

//go:generate mockgen -destination=mocks/mock.go -package=mocks github.com/nkolosov/whip-round/internal/repository User,Session,DB

var (
	ErrDBNil = errors.New("db is nil")
)

type User interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, u *domain.User) (*domain.User, error)
}

type Session interface {
	CreateRefreshToken(ctx context.Context, refreshToken *domain.RefreshSession) error
	GetRefreshToken(ctx context.Context, tokenValue string) (*domain.RefreshSession, error)
	DeleteRefreshToken(ctx context.Context, tokenValue string) error
	DeleteRefreshTokenByUserId(ctx context.Context, userID uuid.UUID) error
}

type DB interface {
	User
	Session
}

type Repository struct {
	User
	Session
}

func NewRepository(db *sql.DB) (*Repository, error) {
	if db == nil {
		return nil, ErrDBNil
	}

	return &Repository{
		Session: NewRepoSessions(db),
		User:    NewUserRepository(db),
	}, nil
}
