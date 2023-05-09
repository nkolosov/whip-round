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
	"sync"
	"testing"
)

func TestNewRepoSessions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockSession(ctrl)

	repo := NewRepoSessions(mockStore)

	if repo == nil {
		t.Errorf("unexpected result, want: not nil, got: nil")
	}
}

func TestCreateRefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockSession(ctrl)
	repo := &SessionsStore{
		db: mockStore,
		mu: sync.RWMutex{},
	}

	//mockStore.EXPECT().CreateRefreshToken(gomock.Any(), gomock.Any()).Return(nil)

	testCases := []struct {
		name           string
		mockError      error
		inputToken     *domain.RefreshSession
		expectedError  error
		expectedCalled bool
	}{
		{
			name:           "success",
			mockError:      nil,
			inputToken:     &domain.RefreshSession{},
			expectedError:  nil,
			expectedCalled: true,
		},
		{
			name:           "error",
			mockError:      errors.New("error"),
			inputToken:     &domain.RefreshSession{},
			expectedError:  errors.New("error"),
			expectedCalled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore.EXPECT().CreateRefreshToken(gomock.Any(), tc.inputToken).Return(tc.mockError).Times(func() int {
				if tc.expectedCalled {
					return 1
				}
				return 0
			}())

			// execute
			err := repo.CreateRefreshToken(context.TODO(), tc.inputToken)

			// verify
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("unexpected error, want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetRefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockSession(ctrl)
	repo := &SessionsStore{
		db: mockStore,
		mu: sync.RWMutex{},
	}

	testCases := []struct {
		name          string
		mockResult    *domain.RefreshSession
		mockErr       error
		inputToken    string
		expectedToken *domain.RefreshSession
		expectedErr   error
	}{
		{
			name:          "success",
			mockResult:    &domain.RefreshSession{},
			mockErr:       nil,
			inputToken:    "token1",
			expectedToken: &domain.RefreshSession{},
			expectedErr:   nil,
		},
		{
			name:          "error",
			mockResult:    nil,
			mockErr:       errors.New("error"),
			inputToken:    "token1",
			expectedToken: nil,
			expectedErr:   errors.New("error"),
		},
		{
			name:          "not found",
			mockResult:    nil,
			mockErr:       memory.ErrSessionNotFound,
			inputToken:    "token1",
			expectedToken: nil,
			expectedErr:   memory.ErrSessionNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore.EXPECT().GetRefreshToken(gomock.Any(), tc.inputToken).Return(tc.mockResult, tc.mockErr)

			token, err := repo.GetRefreshToken(context.TODO(), tc.inputToken)

			if !reflect.DeepEqual(token, tc.expectedToken) {
				t.Errorf("unexpected token, want: %v, got: %v", tc.expectedToken, token)
			}

			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("unexpected error, want: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestDeleteRefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockSession(ctrl)

	repo := &SessionsStore{
		db: mockStore,
		mu: sync.RWMutex{},
	}

	testCases := []struct {
		name          string
		mockError     error
		inputToken    string
		expectedError error
	}{
		{
			name:          "success",
			mockError:     nil,
			inputToken:    "token",
			expectedError: nil,
		},
		{
			name:          "error",
			mockError:     errors.New("error"),
			inputToken:    "token",
			expectedError: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore.EXPECT().DeleteRefreshToken(gomock.Any(), tc.inputToken).Return(tc.mockError)

			// execute
			err := repo.DeleteRefreshToken(context.TODO(), tc.inputToken)

			// verify
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("unexpected error, want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestDeleteRefreshTokenByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockSession(ctrl)
	repo := &SessionsStore{
		db: mockStore,
		mu: sync.RWMutex{},
	}

	userID := uuid.New()
	testCases := []struct {
		name           string
		mockError      error
		expectedError  error
		expectedCalled bool
	}{
		{
			name:           "success",
			mockError:      nil,
			expectedError:  nil,
			expectedCalled: true,
		},
		{
			name:           "error",
			mockError:      errors.New("error"),
			expectedError:  errors.New("error"),
			expectedCalled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore.EXPECT().DeleteRefreshTokenByUserId(gomock.Any(), userID).Return(tc.mockError).Times(func() int {
				if tc.expectedCalled {
					return 1
				}
				return 0
			}())

			// execute
			err := repo.DeleteRefreshTokenByUserId(context.TODO(), userID)

			// verify
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("unexpected error, want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
