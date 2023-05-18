package domain

import (
	"github.com/google/uuid"
	"time"
)

type RefreshSession struct {
	ID        int64
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
}
