package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GopherInn/internal/domain"
	"github.com/saleh-ghazimoradi/GopherInn/internal/dto"
	"github.com/saleh-ghazimoradi/GopherInn/internal/helper"
	"github.com/saleh-ghazimoradi/GopherInn/internal/repository"
	"github.com/saleh-ghazimoradi/GopherInn/utils"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.UserRequest) (*dto.UserResponse, error)
	GetUserById(ctx context.Context, id string) (*dto.UserResponse, error)
	GetUsers(ctx context.Context, page, limit int) ([]*dto.UserResponse, *helper.PaginatedMeta, error)
	UpdateUser(ctx context.Context, id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func (u *userService) CreateUser(ctx context.Context, req *dto.UserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return u.convertDomainToDto(user), nil
}

func (u *userService) GetUserById(ctx context.Context, id string) (*dto.UserResponse, error) {
	user, err := u.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.convertDomainToDto(user), nil
}

func (u *userService) GetUsers(ctx context.Context, page, limit int) ([]*dto.UserResponse, *helper.PaginatedMeta, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	total, err := u.userRepository.CountUsers(ctx)
	if err != nil {
		return nil, nil, err
	}

	users, err := u.userRepository.GetUsers(ctx, offset, limit)
	if err != nil {
		return nil, nil, err
	}

	response := make([]*dto.UserResponse, len(users))

	for i := range users {
		response[i] = u.convertDomainToDto(users[i])
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	meta := &helper.PaginatedMeta{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPages,
	}

	return response, meta, nil
}

func (u *userService) convertDomainToDto(user *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func (u *userService) UpdateUser(ctx context.Context, id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := u.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}

	if req.LastName != nil {
		user.LastName = *req.LastName
	}

	if err = u.userRepository.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return u.convertDomainToDto(user), nil
}

func (u *userService) DeleteUser(ctx context.Context, id string) error {
	return u.userRepository.DeleteUser(ctx, id)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
