package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
)

var (
	ErrUserNotFound = fmt.Errorf("user with email %s not found", "%s")
	ErrUserCreate   = fmt.Errorf("user with email %s not created", "%s")
)

//go:generate mockgen -destination=mocks/mock.go -package=mocks github.com/nkolosov/whip-round/internal/repository User,Session

type UserStore struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (repo *UserStore) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	createUserQuery := `INSERT INTO users (login, email, birthdate, phone, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id uuid.UUID
	err = tx.QueryRowContext(ctx, createUserQuery, u.Login, u.Email, u.Birthdate, u.Phone, u.Balance).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf(ErrUserCreate.Error(), u.Email)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(ErrUserCreate.Error(), u.Email)
	}

	u.ID = id
	return u, nil

}

func (repo *UserStore) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, login, email, birthdate, phone, balance, created_at FROM users WHERE email = $1`
	row := repo.db.QueryRowContext(ctx, query, email)

	user := &domain.User{}
	err := row.Scan(
		&user.ID,
		&user.Login,
		&user.Email,
		&user.Birthdate,
		&user.Phone,
		&user.Balance,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(ErrUserNotFound.Error(), email)
		}
		return nil, err
	}

	return user, nil
}
