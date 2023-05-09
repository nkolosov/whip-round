package domain

import (
	"github.com/nkolosov/whip-round/internal/utils/conv"
	"github.com/nkolosov/whip-round/internal/utils/currency"
)

// ConvertUserDTOToUser converts UserDTO to User domain object.
func ConvertUserDTOToUser(uDTO *UserDTO) (*User, error) {
	if uDTO == nil {
		return nil, ErrEmptyUserDTO
	}

	// convert interface{} to float64
	toFloat64, err := conv.ConvertInterfaceToFloat64(uDTO.Balance)
	if err != nil {
		return nil, err
	}

	// convert dollars to cents
	cents := currency.ConvertDollarsToCents(toFloat64)

	return &User{
		Login:     uDTO.Login,
		Email:     uDTO.Email,
		Birthdate: uDTO.Birthdate,
		Phone:     uDTO.Phone,
		Balance:   cents,
	}, nil
}

// ConvertUserToUserDTO converts User domain object to UserDTO.
func ConvertUserToUserDTO(u *User) *UserDTO {
	if u == nil {
		return nil
	}

	return &UserDTO{
		Login:     u.Login,
		Email:     u.Email,
		Birthdate: u.Birthdate,
		Phone:     u.Phone,
		Balance:   u.Balance,
	}
}
