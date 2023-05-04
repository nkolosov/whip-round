package repository

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/nkolosov/whip-round/internal/repository/mocks"
	"testing"
)

func TestRepository_NewRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authMock := mocks.NewMockUser(ctrl)
	authMock.EXPECT().GetUserByEmail(gomock.Any(), "email").Return(nil, nil)

	sessionMock := mocks.NewMockSession(ctrl)
	sessionMock.EXPECT().GetRefreshToken(gomock.Any(), "tokenValue").Return(nil, nil)

	db := &Repository{
		User:    authMock,
		Session: sessionMock,
	}

	testCases := []struct {
		name          string
		inputDB       DB
		expectedError error
	}{
		{
			name:          "success",
			inputDB:       db,
			expectedError: nil,
		},
		{
			name:          "db is nil",
			inputDB:       nil,
			expectedError: ErrDBNil,
		},
	}

	for _, tc := range testCases {
		repo, err := NewRepository(tc.inputDB)

		if tc.expectedError != nil {
			if err == nil {
				t.Errorf("unexpected error, want: %v, got: nil", tc.expectedError)
			} else if !errors.Is(err, tc.expectedError) {
				t.Errorf("unexpected error, want: %v, got: %v", tc.expectedError, err)
			}
			return
		}

		if err != nil {
			t.Errorf("unexpected error, want: nil, got: %v", err)
			return
		}

		u, err := repo.User.GetUserByEmail(context.TODO(), "email")
		if u != nil || err != nil {
			t.Errorf("unexpected result, want: nil, nil, got: %v, %v", u, err)
		}

		s, err := repo.Session.GetRefreshToken(context.TODO(), "tokenValue")
		if s != nil || err != nil {
			t.Errorf("unexpected result, want: nil, nil, got: %v, %v", s, err)
		}
	}
}
