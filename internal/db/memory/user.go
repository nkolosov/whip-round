package memory

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"strings"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

func (db *DB) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	for _, user := range db.userStore {
		if user.Email == u.Email {
			return nil, ErrUserExists
		}
	}

	u.ID = uuid.New()
	u.Email = strings.ToLower(u.Email)
	db.userStore[u.ID.String()] = u

	fmt.Println("CreateUser: ", u)

	return u, nil
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	for k, user := range db.userStore {
		if user.Email == strings.ToLower(email) {
			fmt.Println("GetUserByEmail: ", user)
			return db.userStore[k], nil
		}
	}

	return nil, ErrUserNotFound
}
