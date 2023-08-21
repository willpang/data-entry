package services

import (
	"context"

	"github.com/willpang/data-entry/models"
)

type UserDatastore interface {
	GetAllUsers(ctx context.Context) []models.User
	GetUser(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, req models.UserCreateRequest) (*models.User, error)
	UpdateUser(ctx context.Context, id string, req models.UserUpdateRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}
