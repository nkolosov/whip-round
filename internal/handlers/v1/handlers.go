package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nkolosov/whip-round/internal/services"
)

type Handlers struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) Init(api *gin.RouterGroup) *gin.RouterGroup {
	usersAPI := api.Group("/users")
	{
		usersAPI.POST("", h.CreateUser)
		usersAPI.GET("", h.GetUserByFilters)
	}

	return api
}
