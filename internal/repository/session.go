package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"sync"
)

type Item struct {
	ID        int64
	UserID    uuid.UUID
	Token     string
	ExpiresAt int64
}

type SessionsStore struct {
	db Session
	mu sync.RWMutex
}

func NewRepoSessions(db Session) *SessionsStore {
	return &SessionsStore{
		db: db,
		mu: sync.RWMutex{},
	}
}

func (repo *SessionsStore) CreateRefreshToken(ctx context.Context, refreshToken *domain.RefreshSession) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	err := repo.db.CreateRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SessionsStore) GetRefreshToken(ctx context.Context, tokenValue string) (*domain.RefreshSession, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	return repo.db.GetRefreshToken(ctx, tokenValue)
}

func (repo *SessionsStore) DeleteRefreshToken(ctx context.Context, tokenValue string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.db.DeleteRefreshToken(ctx, tokenValue)
}

func (repo *SessionsStore) DeleteRefreshTokenByUserId(ctx context.Context, userID uuid.UUID) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.db.DeleteRefreshTokenByUserId(ctx, userID)
}
