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
	db Pool
}

func NewRepoSessions(db Pool) *SessionsStore {
	return &SessionsStore{
		db: db,
	}
}

func (repo *SessionsStore) CreateRefreshToken(ctx context.Context, token *domain.RefreshSession) error {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return err
	}

	createUserQuery := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	err = tx.QueryRow(ctx, createUserQuery, token.UserID, token.Token, token.ExpiresAt).Scan(&id)
	if err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf(ErrSessionNotFound.Error(), token.Token)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf(ErrSessionNotFound.Error(), token.Token)
	}

	token.ID = id
	return nil
}

func (repo *SessionsStore) GetRefreshToken(ctx context.Context, tokenValue string) (*domain.RefreshSession, error) {
	query := `SELECT user_id, token, expires_at FROM refresh_tokens WHERE token = $1`
	row := repo.db.QueryRow(ctx, query, tokenValue)

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
	row := repo.db.QueryRow(ctx, query, tokenValue)

	refreshToken := &domain.RefreshSession{}
	err := row.Scan(
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SessionsStore) DeleteRefreshTokenByUserId(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM refresh_tokens WHERE token = $1`
	row := repo.db.QueryRow(ctx, query, userID)

	refreshToken := &domain.RefreshSession{}
	err := row.Scan(
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
	)
	if err != nil {
		return err
	}

	return nil
}
