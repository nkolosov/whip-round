package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserStore_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	store := NewUserRepository(db)

	// создаем контекст теста
	ctx := context.Background()
	email := "test@example.com"
	expectedUser := &domain.User{ID: uuid.New(), Email: email}

	// создаем тестовые случаи
	testCases := []struct {
		name          string
		rows          *sqlmock.Rows
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name: "success",
			rows: sqlmock.NewRows([]string{"id", "login", "email", "birthdate", "phone", "balance", "created_at"}).
				AddRow(expectedUser.ID, expectedUser.Login, expectedUser.Email, expectedUser.Birthdate, expectedUser.Phone, expectedUser.Balance, expectedUser.CreatedAt),
			expectedUser:  expectedUser,
			expectedError: nil,
		},
		{
			name:          "not found",
			rows:          sqlmock.NewRows([]string{"id", "login", "email", "birthdate", "phone", "balance", "created_at"}),
			expectedUser:  nil,
			expectedError: fmt.Errorf(ErrUserNotFound.Error(), email),
		},
	}

	// выполняем тестовые случаи
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// устанавливаем мок ожидания
			mock.ExpectQuery("^SELECT (.+) FROM users WHERE email = \\$1").
				WithArgs(email).
				WillReturnRows(tc.rows)

			// вызываем метод GetUser
			user, err := store.GetUserByEmail(ctx, email)

			// проверяем результаты
			require.Equal(t, tc.expectedUser, user)
			require.Equal(t, tc.expectedError, err)

			// убеждаемся, что все ожидания мока были выполнены
			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUserStore_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	userRepo := NewUserRepository(db)

	// Создаем контекст теста
	ctx := context.Background()
	user := &domain.User{
		Login:     "test",
		Email:     "test@example.com",
		Birthdate: "01.01.2000",
		Phone:     "123456789",
		Balance:   0,
		CreatedAt: time.Now(),
	}
	expectedUser := &domain.User{
		ID:        uuid.New(),
		Login:     user.Login,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Phone:     user.Phone,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}

	// Создаем тестовые случаи
	testCases := []struct {
		name          string
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "success",
			expectedUser:  expectedUser,
			expectedError: nil,
		},
		{
			name:          "error",
			expectedUser:  nil,
			expectedError: fmt.Errorf("пользователь с адресом электронной почты %s не создан", user.Email),
		},
	}

	// Выполняем тестовые случаи
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedError != nil {
				mock.ExpectBegin().WillReturnError(tc.expectedError)
			} else {
				mock.ExpectBegin()
				mock.ExpectQuery("^INSERT INTO users (.+) VALUES (.+) RETURNING (.+)").
					WithArgs(user.Login, user.Email, user.Birthdate, user.Phone, user.Balance).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(expectedUser.ID))
				mock.ExpectCommit()
			}

			// вызываем метод CreateUser
			createdUser, err := userRepo.CreateUser(ctx, user)

			// проверяем результаты
			require.Equal(t, tc.expectedUser, createdUser)
			require.Equal(t, tc.expectedError, err)

			// убеждаемся, что все ожидания мока были выполнены
			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}
