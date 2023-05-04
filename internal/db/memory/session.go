package memory

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

func (db *DB) CreateRefreshToken(ctx context.Context, refreshToken *domain.RefreshSession) error {
	db.sessionStore = append(db.sessionStore, refreshToken)
	return nil
}

func (db *DB) GetRefreshToken(ctx context.Context, tokenValue string) (*domain.RefreshSession, error) {
	for _, session := range db.sessionStore {
		if session.Token == tokenValue {
			return session, nil
		}
	}

	return nil, ErrSessionNotFound
}

func (db *DB) DeleteRefreshToken(ctx context.Context, tokenValue string) error {
	for k, session := range db.sessionStore {
		if session.Token == tokenValue {
			db.sessionStore = append(db.sessionStore[:k], db.sessionStore[k+1:]...)
			return nil
		}
	}

	return ErrSessionNotFound
}

func (db *DB) DeleteRefreshTokenByUserId(ctx context.Context, userID uuid.UUID) error {
	for k, session := range db.sessionStore {
		if session.UserID == userID {
			db.sessionStore = append(db.sessionStore[:k], db.sessionStore[k+1:]...)
			return nil
		}
	}

	return ErrSessionNotFound
}
