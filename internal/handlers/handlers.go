package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nkolosov/whip-round/docs"
	v1 "github.com/nkolosov/whip-round/internal/handlers/v1"
	"github.com/nkolosov/whip-round/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	service *services.Service
}

// NewHandler creates a new Handlers with the necessary dependencies.
func NewHandler(service *services.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

// Init initializes the rest transport handlers and returns a gin engine.
func (h *Handlers) Init(router *gin.Engine) *gin.Engine {
	//router := gin.Default()
	//router.Use(sessions.SessionStore("session", h.store))

	// Init swagger
	initSwagger(router)

	apiV1 := router.Group("/api/v1")
	{
		handlersV1 := v1.NewHandler(h.service)
		handlersV1.Init(apiV1)
	}

	return router
}

// @title Your API Title
// @description Your API Description
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func initSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
