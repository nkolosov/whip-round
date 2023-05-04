package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck is a handler that returns a 200 status code to indicate that the server is running.
// @Summary Check server health
// @Description Check server health
// @ID health-check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /health-check [get]
func (h *Handlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
