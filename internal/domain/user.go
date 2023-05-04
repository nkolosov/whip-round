package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/nkolosov/whip-round/internal/utils/conv"
	"github.com/nkolosov/whip-round/internal/utils/currency"
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

type UserDTO struct {
	Login     string      `json:"login" example:"john.doe" swaggertype:"string"`
	Email     string      `json:"email" example:"jogn.doe@mail.com" swaggertype:"string"`
	Birthdate string      `json:"birthdate" example:"1990-01-01" format:"date" swaggertype:"string"`
	Phone     string      `json:"phone" example:"+79999999999" swaggertype:"string"`
	Balance   interface{} `json:"balance" example:"100.00" swaggertype:"number,string"` // in dollars
}

func (cdto *UserDTO) FromDomain(c *User) {
	if c == nil {
		return
	}

	cdto.Login = c.Login
	cdto.Email = c.Email
	cdto.Birthdate = c.Birthdate
	cdto.Phone = c.Phone
	cdto.Balance = c.Balance
}

func (cdto *UserDTO) ToDomain() (*User, error) {
	if cdto == nil {
		return nil, ErrEmptyUserDTO
	}

	// convert interface{} to float64
	toFloat64, err := conv.ConvertInterfaceToFloat64(cdto.Balance)
	if err != nil {
		return nil, err
	}

	// convert dollars to cents
	cents := currency.ConvertDollarsToCents(toFloat64)

	return &User{
		Login:     cdto.Login,
		Email:     cdto.Email,
		Birthdate: cdto.Birthdate,
		Phone:     cdto.Phone,
		Balance:   cents,
	}, nil
}

func (cdto *UserDTO) Validate() error {
	if cdto == nil {
		return ErrEmptyUserDTO
	}

	return nil
}
