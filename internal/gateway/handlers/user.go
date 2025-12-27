package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/saleh-ghazimoradi/GopherInn/internal/dto"
	"github.com/saleh-ghazimoradi/GopherInn/internal/helper"
	"github.com/saleh-ghazimoradi/GopherInn/internal/repository"
	"github.com/saleh-ghazimoradi/GopherInn/internal/service"
	"strconv"
)

type UserHandler struct {
	userService service.UserService
	logger      *zerolog.Logger
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var payload dto.UserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		h.logger.Warn().Err(err).Msg("failed to bind payload")
		helper.BadRequestResponse(ctx, "given invalid user request", err)
		return
	}

	user, err := h.userService.CreateUser(ctx, &payload)
	if err != nil {
		h.logger.Err(err).Msg("failed to create user")
		helper.InternalServerError(ctx, "user creation failed", err)
		return
	}

	helper.CreatedResponse(ctx, "user successfully registered", user)
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.userService.GetUserById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			h.logger.Err(err).Msg("user not found")
			helper.NotFoundResponse(ctx, "user not found")
		default:
			h.logger.Err(err).Msg("failed to get user by id")
			helper.InternalServerError(ctx, "failed to get user by id", err)
		}
		return
	}

	helper.SuccessResponse(ctx, "user successfully retrieved", user)
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	users, meta, err := h.userService.GetUsers(ctx, page, limit)
	if err != nil {
		h.logger.Err(err).Msg("failed to get users")
		helper.InternalServerError(ctx, "failed to get users", err)
		return
	}

	helper.PaginatedSuccessResponse(ctx, "users successfully retrieved", users, *meta)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var payload dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		h.logger.Warn().Err(err).Msg("failed to bind payload")
		helper.BadRequestResponse(ctx, "failed to bind payload", err)
		return
	}

	updatedUser, err := h.userService.UpdateUser(ctx, id, &payload)
	if err != nil {
		h.logger.Err(err).Msg("failed to update user")
		helper.InternalServerError(ctx, "failed to update user", err)
		return
	}

	helper.SuccessResponse(ctx, "user successfully updated", updatedUser)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.userService.DeleteUser(ctx, id); err != nil {
		h.logger.Err(err).Msg("failed to delete user")
		helper.InternalServerError(ctx, "failed to delete user", err)
		return
	}

	helper.SuccessResponse(ctx, "user successfully deleted", nil)
}

func NewUserHandler(userService service.UserService, logger *zerolog.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}
