package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nkolosov/whip-round/internal/domain"
	"github.com/nkolosov/whip-round/internal/utils/currency"
)

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user
// @ID create-user
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRequest true "UserRequest"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *Handlers) CreateUser(c *gin.Context) {
	var userDTO *domain.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	createUser, err := h.service.UserService.CreateUser(c, userDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to create user: %s with error %s", userDTO, err.Error()))
		return
	}

	createUser.Balance = currency.ConvertDollarsToCents(float64(createUser.Balance))

	c.JSON(http.StatusOK, UserResponse{
		ID:        createUser.ID,
		Login:     createUser.Login,
		Email:     createUser.Email,
		Birthdate: createUser.Birthdate,
		Phone:     createUser.Phone,
		CreatedAt: createUser.CreatedAt,
	})
}

// GetUserByFilters gets a user by filters.
// @Summary Get a user by filters
// @Description Get a user by filters
// @ID get-user
// @Tags users
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *Handlers) GetUserByFilters(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		newErrorResponse(c, http.StatusBadRequest, "email parameter is required")
		return
	}

	customer, err := h.service.UserService.FindUserByEmail(c, email)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("failed to find user by email: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        customer.ID,
		Login:     customer.Login,
		Email:     customer.Email,
		Birthdate: customer.Birthdate,
		Phone:     customer.Phone,
		CreatedAt: customer.CreatedAt,
	})
}
