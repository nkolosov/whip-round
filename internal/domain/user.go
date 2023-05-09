package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	ErrEmptyUserDTO = errors.New("empty user dto")
)

type User struct {
	ID        uuid.UUID `json:"id" example:"1" format:"int64" swaggertype:"integer"`
	Login     string    `json:"login" example:"john.doe" swaggertype:"string"`
	Email     string    `json:"email" example:"jogn.doe@mail.com" swaggertype:"string"`
	Birthdate string    `json:"birthdate" example:"1990-01-01" format:"date" swaggertype:"string"`
	Phone     string    `json:"phone" example:"+79999999999" swaggertype:"string"`
	Balance   int64     `json:"balance" example:"100" format:"int64" swaggertype:"integer"` // in cents
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z" swaggertype:"string"`
}

func (u *User) String() string {
	if u == nil {
		return "User(nil)"
	}

	return fmt.Sprintf(
		"User{ID: %s, Login: %s, Email: %s, Birthdate: %s, Phone: %s, Balance: %d, CreatedAt: %s}",
		u.ID, u.Login, u.Email, u.Birthdate, u.Phone, u.Balance, u.CreatedAt,
	)
}

type UserDTO struct {
	Login     string      `json:"login" example:"john.doe" swaggertype:"string"`
	Email     string      `json:"email" example:"jogn.doe@mail.com" swaggertype:"string"`
	Birthdate string      `json:"birthdate" example:"1990-01-01" format:"date" swaggertype:"string"`
	Phone     string      `json:"phone" example:"+79999999999" swaggertype:"string"`
	Balance   interface{} `json:"balance" example:"100.00" swaggertype:"number,string"` // in dollars
}

func (uDTO *UserDTO) String() string {
	if uDTO == nil {
		return "UserDTO(nil)"
	}

	return fmt.Sprintf(
		"UserDTO{Login: %s, Email: %s, Birthdate: %s, Phone: %s, Balance: %s}",
		uDTO.Login, uDTO.Email, uDTO.Birthdate, uDTO.Phone, uDTO.Balance,
	)
}

func (uDTO *UserDTO) Validate() error {
	if uDTO == nil {
		return ErrEmptyUserDTO
	}

	return nil
}
