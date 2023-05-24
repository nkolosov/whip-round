package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/nkolosov/whip-round/internal/domain"
)

//go:generate mockgen -destination=mocks/mock.go -package=mocks github.com/nkolosov/whip-round/internal/repository User,Session,DB
//go:generate mockgen -destination=mocks/mock_pool.go -package=mocks github.com/nkolosov/whip-round/internal/repository Pool,Row
//go:generate mockgen -destination=mocks/mock_tx.go -package=mocks github.com/jackc/pgx/v4 Tx

var (
	ErrDBNil = errors.New("db is nil")
)

type Pool interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}
type Row interface {
	Scan(dest ...interface{}) error
}

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

func NewRepository(db Pool) (*Repository, error) {
	if db == nil {
		return nil, ErrDBNil
	}

	return &Repository{
		Session: NewRepoSessions(db),
		User:    NewUserRepository(db),
	}, nil
}
