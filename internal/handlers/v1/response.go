package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ErrorResponse is the error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// newErrorResponse creates a new error response and logs the error message
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

type UserRequestFilters struct {
	Email string `json:"email,omitempty" example:"some@mail.com" swaggertype:"string"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id" example:"1" format:"int64" swaggertype:"integer"`
	Email     string    `json:"email" example:"john.doe@example.com" swaggertype:"string"`
	Password  string    `json:"password,omitempty" example:"password" swaggertype:"string"`
	Name      string    `json:"name,omitempty" example:"John Doe" swaggertype:"string"`
	Login     string    `json:"login" example:"john.doe" swaggertype:"string"`
	Birthdate string    `json:"birthdate" example:"1990-01-01" format:"date" swaggertype:"string"`
	Phone     string    `json:"phone" example:"+79999999999" swaggertype:"string"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z" swaggertype:"string"`
}

type UserRequest struct {
	Login     string      `json:"login" example:"john.doe" swaggertype:"string"`
	Email     string      `json:"email" example:"john.doe@example.com" swaggertype:"string"`
	Birthdate string      `json:"birthdate" example:"1990-01-01" format:"date" swaggertype:"string"`
	Phone     string      `json:"phone" example:"+79999999999" swaggertype:"string"`
	Balance   interface{} `json:"balance" example:"100.00" swaggertype:"number"`
}
