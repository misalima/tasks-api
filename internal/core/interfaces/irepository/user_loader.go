package irepository

import (
	"context"
	"tasks-api/internal/core/domain"
)

type UserLoader interface {
	InsertUser(ctx context.Context, user domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}
