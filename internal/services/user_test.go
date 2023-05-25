package services

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nkolosov/whip-round/internal/repository"

	//"github.com/nkolosov/whip-round/internal/db/memory"
	"github.com/nkolosov/whip-round/internal/domain"
	mockshash "github.com/nkolosov/whip-round/internal/hash/mocks"
	mocksrepo "github.com/nkolosov/whip-round/internal/repository/mocks"
	mockstoken "github.com/nkolosov/whip-round/internal/token/mocks"
)

func TestUserService_FindUserByEmailPassword(t *testing.T) {
	testCases := []struct {
		name           string
		email          string
		password       string
		mockResultUser *domain.User
		mockResultErr  error
		expectedResult *domain.User
		expectedError  error
	}{
		{
			name:     "valid credentials",
			email:    "test@example.com",
			password: "hashed_password",
			mockResultUser: &domain.User{
				Email: "test@example.com",
			},
			mockResultErr: nil,
			expectedResult: &domain.User{
				Email: "test@example.com",
			},
			expectedError: nil,
		},
		{
			name:           "invalid credentials",
			email:          "non_existing@example.com",
			password:       "non_existing_password",
			mockResultUser: nil,
			mockResultErr:  fmt.Errorf(repository.ErrUserNotFound.Error(), "non_existing@example.com"),
			expectedResult: nil,
			expectedError:  fmt.Errorf(repository.ErrUserNotFound.Error(), "non_existing@example.com"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			userRepo := mocksrepo.NewMockUser(ctrl)
			tokenManager := mockstoken.NewMockManager(ctrl)
			hashManager := mockshash.NewMockHashManager(ctrl)

			userService := NewUserService(userRepo, tokenManager, hashManager)

			userRepo.EXPECT().GetUserByEmail(ctx, tc.email).Return(tc.mockResultUser, tc.mockResultErr)

			user, err := userService.FindUserByEmail(ctx, tc.email)
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(user, tc.expectedResult) {
				t.Errorf("Expected %v, got %v", tc.expectedResult, user)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userRepo := mocksrepo.NewMockUser(ctrl)
	hashManager := mockshash.NewMockHashManager(ctrl)
	tokenManager := mockstoken.NewMockManager(ctrl)

	userService := NewUserService(userRepo, tokenManager, hashManager)

	existingUser := &domain.User{
		Email:     "test@example.com",
		Birthdate: "1990-01-01",
		Login:     "test",
		Phone:     "123456789",
		Balance:   123123,
	}
	userRepo.EXPECT().GetUserByEmail(ctx, existingUser.Email).Return(existingUser, nil)

	validUser := &domain.User{
		Email:     "new_user@example.com",
		Birthdate: "1990-01-01",
		Login:     "test",
		Phone:     "123456789",
		Balance:   21342134,
	}

	t.Run("User with the same email already exists", func(t *testing.T) {
		user := existingUser
		var expectedUser *domain.User = nil
		expectedError := fmt.Errorf(ErrUserAlreadyExists.Error(), "test@example.com")

		user, err := userService.CreateUser(ctx, &domain.UserDTO{
			Email: user.Email,
		})

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}

		if !reflect.DeepEqual(user, expectedUser) {
			t.Errorf("Expected user %v, got %v", expectedUser, user)
		}
	})

	t.Run("Valid user", func(t *testing.T) {
		user := validUser
		expectedUser := validUser
		expectedError := error(nil)

		//hashedPassword := "hashedPassword"
		//hashManager.EXPECT().HashPassword(user.Password).Return(hashedPassword, nil)
		userRepo.EXPECT().GetUserByEmail(ctx, user.Email).Return(nil, errors.New("user not found"))
		userRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(expectedUser, nil)

		user, err := userService.CreateUser(ctx, &domain.UserDTO{
			Email:     user.Email,
			Birthdate: "1990-01-01",
			Login:     "test",
			Phone:     "123456789",
			Balance:   21342134,
		})

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}

		if !reflect.DeepEqual(user, expectedUser) {
			t.Errorf("Expected user %v, got %v", expectedUser, user)
		}
	})
}
