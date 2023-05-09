package token

import (
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/utils/random"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	userID := uuid.New()
	pasetoMaker, err := NewPasetoManager(random.RandomString(32))
	if err != nil {
		t.Fatal(err)
	}

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := pasetoMaker.CreateToken(userID, duration)
	if err != nil {
		t.Fatal(err)
	}

	payload, err := pasetoMaker.VerifyToken(token)
	if err != nil {
		t.Fatal(err)
	}

	if payload.UserID != userID {
		t.Fatal("user_id is not correct")
	}
	if payload.IssuedAt < time.Duration(issuedAt.Unix()) {
		t.Fatal("issued_at is not correct")
	}
	if payload.ExpiredAt > time.Duration(expiredAt.Unix()) {
		t.Fatal("expired_at is not correct")
	}
}

func TestExpiredPasetoToken(t *testing.T) {
	pasetoMaker, err := NewPasetoManager(random.RandomString(32))
	if err != nil {
		t.Fatal(err)
	}

	wrongDuration := -time.Minute
	token, err := pasetoMaker.CreateToken(uuid.New(), wrongDuration)
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("token is empty")
	}

	payload, err := pasetoMaker.VerifyToken(token)
	if err != ErrExpiredToken {
		t.Fatal("token is not expired")
	}
	if payload != nil {
		t.Fatal("payload is not nil")
	}
}
