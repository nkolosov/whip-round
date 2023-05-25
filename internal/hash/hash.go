package hash

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -destination=mocks/mock_hash.go -package=mocks github.com/nkolosov/whip-round/internal/hash HashManager
type HashManager interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

var (
	ErrHash = fmt.Errorf("error hashing")
)

type Manager struct {
	salt []byte
}

func NewHash(salt string) (*Manager, error) {
	if salt == "" {
		return &Manager{}, ErrHash
	}

	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return &Manager{}, fmt.Errorf("error decoding salt: %w", err)
	}

	return &Manager{salt: saltBytes}, nil
}

func (h *Manager) HashPassword(password string) (string, error) {
	saltedPassword := append([]byte(password), h.salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHash
	}

	return string(hashedPassword), nil
}

func (h *Manager) CheckPasswordHash(password, hash string) bool {
	saltedPassword := append([]byte(password), h.salt...)
	err := bcrypt.CompareHashAndPassword([]byte(hash), saltedPassword)
	return err == nil
}
