package repository

import (
	"context"
	"github.com/nkolosov/whip-round/internal/db/memory"
	"github.com/nkolosov/whip-round/internal/domain"
	"sync"
)

type UserStore struct {
	store memory.UserDB
	mu    sync.RWMutex
}

func NewUserRepository(store memory.UserDB) *UserStore {
	return &UserStore{
		store: store,
		mu:    sync.RWMutex{},
	}
}

func (repo *UserStore) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.store.CreateUser(ctx, u)
}

func (repo *UserStore) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	return repo.store.GetUserByEmail(ctx, email)
}
