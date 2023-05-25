package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"github.com/nkolosov/whip-round/internal/repository/mocks"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserStore_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock Pool and Row
	mockPool := mocks.NewMockPool(ctrl)
	mockRow := mocks.NewMockRow(ctrl)

	store := NewUserRepository(mockPool)

	// создаем контекст теста
	ctx := context.Background()
	email := "test@example.com"
	expectedUser := &domain.User{
		ID:        uuid.New(),
		Email:     email,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPool.EXPECT().QueryRow(ctx, gomock.Any(), email).Return(mockRow)
			mockRow.EXPECT().Scan(gomock.Any()).DoAndReturn(func(dest ...interface{}) error {
				if tc.expectedUser == nil {
					return sql.ErrNoRows // simulate no rows found
				}
				for i, d := range dest {
					switch v := d.(type) {
					case *uuid.UUID:
						*v = expectedUser.ID
					case *string:
						if i == 2 { // index of email in the scan
							*v = expectedUser.Email
						}
						// handle other types here
					}
				}
				return nil
			})

			user, err := store.GetUserByEmail(ctx, email)

			// проверяем результаты
			require.Equal(t, tc.expectedUser, user)
			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserStore_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTx(ctrl)
	mockRow := mocks.NewMockRow(ctrl)

	store := NewUserRepository(mockTx)

	// Тестовые случаи
	testCases := []struct {
		name          string
		user          *domain.User
		expectedError error
	}{
		{
			name: "success",
			user: &domain.User{
				Login:     "test",
				Email:     "test@example.com",
				Birthdate: "01.01.2000",
				Phone:     "123456789",
				Balance:   0,
				CreatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name: "error",
			user: &domain.User{
				Login:     "test",
				Email:     "test@example.com",
				Birthdate: "01.01.2000",
				Phone:     "123456789",
				Balance:   0,
				CreatedAt: time.Now(),
			},
			expectedError: fmt.Errorf(ErrUserCreate.Error(), "test@example.com"),
		},
	}

	// Выполняем тестовые случаи
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTx.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)

			if tc.expectedError == nil {
				mockTx.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(nil)
				mockTx.EXPECT().Commit(gomock.Any()).Return(nil)
			} else {
				mockTx.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(tc.expectedError)
				mockTx.EXPECT().Rollback(gomock.Any()).Return(nil)
			}

			createdUser, err := store.CreateUser(context.Background(), tc.user)

			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotNil(t, createdUser)
			} else {
				require.Nil(t, createdUser)
			}
		})
	}
}
