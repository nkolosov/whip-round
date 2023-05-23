package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nkolosov/whip-round/internal/domain"
	mockshash "github.com/nkolosov/whip-round/internal/hash/mocks"
	mocksrepo "github.com/nkolosov/whip-round/internal/repository/mocks"
	mockstoken "github.com/nkolosov/whip-round/internal/token/mocks"
)

func TestUserService_GetAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userRepo := mocksrepo.NewMockUser(ctrl)
	hashManager := mockshash.NewMockHashManager(ctrl)
	tokenManager := mockstoken.NewMockManager(ctrl)

	userService := NewAuthService(userRepo, tokenManager, hashManager)

	email := "test@example.com"
	password := "password"
	//hashedPassword := "hashed_password"

	user := &domain.User{
		ID:    uuid.New(),
		Email: email,
	}
	userRepo.EXPECT().GetUserByEmail(ctx, email).Return(user, nil)

	expectedAccessToken := "access_token"
	expectedRefreshToken := "refresh_token"
	tokenManager.EXPECT().CreateToken(user.ID, time.Hour*24).Return(expectedAccessToken, nil)
	tokenManager.EXPECT().CreateToken(user.ID, time.Hour*24*7).Return(expectedRefreshToken, nil)

	t.Run("Valid sign in", func(t *testing.T) {

		//hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)

		gotAccessToken, gotRefreshToken, err := userService.GetAccessToken(ctx, email, password)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if gotAccessToken != expectedAccessToken {
			t.Errorf("Expected access token %v, got %v", expectedAccessToken, gotAccessToken)
		}

		if gotRefreshToken != expectedRefreshToken {
			t.Errorf("Expected refresh token %v, got %v", expectedRefreshToken, gotRefreshToken)
		}
	})

	t.Run("Error getting user by email", func(t *testing.T) {
		expectedAccessToken := ""
		expectedRefreshToken := ""
		expectedError := errors.New("database error")

		userRepo.EXPECT().GetUserByEmail(ctx, email).Return(nil, expectedError)

		gotAccessToken, gotRefreshToken, err := userService.GetAccessToken(ctx, email, password)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}

		if gotAccessToken != expectedAccessToken {
			t.Errorf("Expected access token %v, got %v", expectedAccessToken, gotAccessToken)
		}

		if gotRefreshToken != expectedRefreshToken {
			t.Errorf("Expected refresh token %v, got %v", expectedRefreshToken, gotRefreshToken)
		}
	})
}
