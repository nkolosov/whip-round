package token

import (
	"time"

	"github.com/google/uuid"
)

// Manager is an interface for managing token.
//
//go:generate mockgen -destination=mocks/mocks.go -package=mocks github.com/nkolosov/whip-round/internal/token Manager
type Manager interface {
	// CreateToken creates a new token for specific username and duration.
	CreateToken(userID uuid.UUID, duration time.Duration) (string, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
