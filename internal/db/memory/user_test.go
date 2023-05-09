package memory

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"reflect"
	"testing"
)

func TestDB_CreateUser(t *testing.T) {
	duplicateEmail := "john.doe@example.com"

	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	testCases := []struct {
		name        string
		user        *domain.User
		expectedErr error
	}{
		{
			name: "Create user with unique email",
			user: &domain.User{
				ID:    uuid.New(),
				Email: "22john.doe@example.com",
			},
			expectedErr: nil,
		},
		{
			name: "Create user with duplicate email",
			user: &domain.User{
				ID:    uuid.New(),
				Email: duplicateEmail,
			},
			expectedErr: ErrUserExists,
		},
	}

	// Create initial user for testing duplicate email
	initialUser := &domain.User{
		ID:    uuid.New(),
		Email: duplicateEmail,
	}
	db.userStore[initialUser.ID.String()] = initialUser

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err = db.CreateUser(context.Background(), tc.user)
			if err != tc.expectedErr {
				t.Errorf("expected error: %v, but got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestDB_GetUserByEmail(t *testing.T) {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	// Create test users
	user1 := &domain.User{
		ID:    uuid.New(),
		Email: "user1@example.com",
	}
	db.userStore[user1.ID.String()] = user1

	user2 := &domain.User{
		ID:    uuid.New(),
		Email: "user2@example.com",
	}
	db.userStore[user2.ID.String()] = user2

	// Test cases
	tests := []struct {
		name          string
		email         string
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "get existing user by email",
			email:         "user1@example.com",
			expectedUser:  user1,
			expectedError: nil,
		},
		{
			name:          "get existing user by email, case insensitive",
			email:         "USER2@example.com",
			expectedUser:  user2,
			expectedError: nil,
		},
		{
			name:          "get non-existing user by email",
			email:         "user3@example.com",
			expectedUser:  nil,
			expectedError: errors.New("user not found"),
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualUser, actualError := db.GetUserByEmail(context.Background(), tt.email)

			if !reflect.DeepEqual(actualUser, tt.expectedUser) {
				t.Errorf("expected user %v, got %v", tt.expectedUser, actualUser)
			}

			if !reflect.DeepEqual(actualError, tt.expectedError) {
				t.Errorf("expected error %v, got %v", tt.expectedError, actualError)
			}
		})
	}
}
