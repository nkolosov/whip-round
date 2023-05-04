package repository

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/db/memory"
	"github.com/nkolosov/whip-round/internal/domain"
	"github.com/nkolosov/whip-round/internal/repository/mocks"
	"reflect"
	"testing"
)

func TestUserStore_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// create mock UserDB
	mockDB := mocks.NewMockUser(ctrl)

	// create UserStore
	userStore := &UserStore{store: mockDB}

	// create test context
	ctx := context.Background()
	email := "test@example.com"
	expectedUser := &domain.User{ID: uuid.New(), Email: email}

	// create test cases
	testCases := []struct {
		name          string
		dbResponse    *domain.User
		dbError       error
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "success",
			dbResponse:    expectedUser,
			dbError:       nil,
			expectedUser:  expectedUser,
			expectedError: nil,
		},
		{
			name:          "not found",
			dbResponse:    nil,
			dbError:       memory.ErrUserNotFound,
			expectedUser:  nil,
			expectedError: memory.ErrUserNotFound,
		},
		{
			name:          "db error",
			dbResponse:    nil,
			dbError:       errors.New("db error"),
			expectedUser:  nil,
			expectedError: errors.New("db error"),
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// set mock expectations
			mockDB.EXPECT().GetUserByEmail(ctx, email).Return(tc.dbResponse, tc.dbError)

			// call the GetUser method
			user, err := userStore.GetUserByEmail(ctx, email)

			// check the results
			if !reflect.DeepEqual(user, tc.expectedUser) {
				t.Errorf("unexpected user, expected %v, got %v", tc.expectedUser, user)
			}

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("unexpected error, expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestUserStore_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockDB(ctrl)
	store := NewUserRepository(mockStore)

	testCases := []struct {
		name          string
		user          *domain.User
		mockError     error
		expectedUser  *domain.User
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
			},
			mockError:     nil,
			expectedUser:  &domain.User{ID: uuid.New()},
			expectedError: nil,
		},
		{
			name: "error",
			user: &domain.User{
				Login:     "test2",
				Email:     "test2@example.com",
				Birthdate: "01.01.2000",
				Phone:     "123456789",
				Balance:   0,
			},
			mockError:     errors.New("create user error"),
			expectedUser:  nil,
			expectedError: errors.New("create user error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore.EXPECT().CreateUser(gomock.Any(), tc.user).Return(tc.expectedUser, tc.mockError)

			user, err := store.CreateUser(context.Background(), tc.user)

			if !reflect.DeepEqual(user, tc.expectedUser) {
				t.Errorf("unexpected user, want: %v, got: %v", tc.expectedUser, user)
			}

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("unexpected error, want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
