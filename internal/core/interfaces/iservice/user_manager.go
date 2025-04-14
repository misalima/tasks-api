package iservice

import (
	"context"
	"tasks-api/internal/core/domain"
)

type UserManager interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
}
