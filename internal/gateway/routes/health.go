package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saleh-ghazimoradi/GopherInn/internal/gateway/handlers"
)

type HealthRoutes struct {
	healthHandler *handlers.HealthHandler
}

func (h *HealthRoutes) HealthRoute(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/health", h.healthHandler.HealthHandler)
}

func NewHealthRoutes(healthHandler *handlers.HealthHandler) *HealthRoutes {
	return &HealthRoutes{
		healthHandler: healthHandler,
	}
}
