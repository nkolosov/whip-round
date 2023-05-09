package memory

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/domain"
	"testing"
)

func TestDB_CreateRefreshToken(t *testing.T) {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	session := &domain.RefreshSession{
		UserID: uuid.New(),
		Token:  "test-token",
	}

	err = db.CreateRefreshToken(context.Background(), session)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(db.sessionStore) != 1 {
		t.Errorf("expected session store length to be 1, got %d", len(db.sessionStore))
	}

	if db.sessionStore[0] != session {
		t.Errorf("expected session to be added to store")
	}
}

func TestDB_GetRefreshToken(t *testing.T) {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	token := "test-token"
	session := &domain.RefreshSession{
		UserID: uuid.New(),
		Token:  token,
	}

	db.sessionStore = []*domain.RefreshSession{session}

	got, err := db.GetRefreshToken(context.Background(), token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got != session {
		t.Errorf("expected session to be retrieved, got %+v", got)
	}
}

func TestDB_DeleteRefreshToken(t *testing.T) {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	token := "test-token"
	session := &domain.RefreshSession{
		UserID: uuid.New(),
		Token:  token,
	}

	db.sessionStore = []*domain.RefreshSession{session}

	err = db.DeleteRefreshToken(context.Background(), token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(db.sessionStore) != 0 {
		t.Errorf("expected session store length to be 0, got %d", len(db.sessionStore))
	}

	// Try to delete session again - should return error
	err = db.DeleteRefreshToken(context.Background(), token)
	if !errors.Is(err, ErrSessionNotFound) {
		t.Errorf("expected ErrSessionNotFound, got %v", err)
	}
}

func TestDB_GetRefreshToken_ErrSessionNotFound(t *testing.T) {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	// Add a session to the store
	token := "test-token"
	session := &domain.RefreshSession{
		UserID: uuid.New(),
		Token:  token,
	}
	db.sessionStore = append(db.sessionStore, session)

	// Try to get a session with a non-existent token
	tokenValue := "non-existent-token"
	_, err = db.GetRefreshToken(context.Background(), tokenValue)

	// Check that the error is of type ErrSessionNotFound
	if !errors.Is(err, ErrSessionNotFound) {
		t.Errorf("expected ErrSessionNotFound, got %v", err)
	}
}

func TestDB_DeleteRefreshTokenByUserID(t *testing.T) {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("failed to create memory DB: %v", err)
	}

	userID := uuid.New()
	token1 := "test-token-1"
	session1 := &domain.RefreshSession{
		UserID: userID,
		Token:  token1,
	}

	db.sessionStore = []*domain.RefreshSession{session1}

	err = db.DeleteRefreshTokenByUserId(context.Background(), userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	fmt.Println(db.sessionStore)

	if len(db.sessionStore) != 0 {
		t.Errorf("expected session store length to be 0, got %d", len(db.sessionStore))
	}

	// Try to delete sessions again - should return error
	err = db.DeleteRefreshTokenByUserId(context.Background(), userID)
	if !errors.Is(err, ErrSessionNotFound) {
		t.Errorf("expected ErrSessionNotFound, got %v", err)
	}
}
