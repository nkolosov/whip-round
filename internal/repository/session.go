package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
)

type Item struct {
	ID        int64
	UserID    uuid.UUID
	Token     string
	ExpiresAt int64
}

var (
	ErrSessionNotFound = fmt.Errorf("session with token %s not found", "%s")
)

type SessionsStore struct {
	db *sql.DB
}

func NewRepoSessions(db *sql.DB) *SessionsStore {
	return &SessionsStore{
		db: db,
	}
}

func (repo *SessionsStore) CreateRefreshToken(ctx context.Context, token *domain.RefreshSession) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO refresh_tokens (user_id, token, expires_at) values ($1, $2, $3)", token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (repo *SessionsStore) GetRefreshToken(ctx context.Context, tokenValue string) (*domain.RefreshSession, error) {
	query := `SELECT user_id, token, expires_at FROM refresh_tokens WHERE token = $1`
	row := repo.db.QueryRowContext(ctx, query, tokenValue)

	refreshToken := &domain.RefreshSession{}
	err := row.Scan(
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(ErrSessionNotFound.Error(), tokenValue)
		}

		return nil, err
	}

	return refreshToken, nil
}

func (repo *SessionsStore) DeleteRefreshToken(ctx context.Context, tokenValue string) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	_, err := repo.db.ExecContext(ctx, query, tokenValue)
	return err
}

func (repo *SessionsStore) DeleteRefreshTokenByUserId(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM refresh_tokens WHERE token = $1`
	_, err := repo.db.ExecContext(ctx, query, userID)
	return err
}
