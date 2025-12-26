package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/saleh-ghazimoradi/GopherInn/internal/helper"
)

type HealthHandler struct{}

func (h *HealthHandler) HealthHandler(ctx *gin.Context) {
	helper.SuccessResponse(ctx, "I'm breathing", nil)
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
