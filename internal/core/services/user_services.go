package services

import (
	"context"
	"github.com/google/uuid"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/irepository"
	"tasks-api/internal/core/interfaces/iservice"
	"time"
)

var _ iservice.UserManager = (*UserServices)(nil)

type UserServices struct {
	repo irepository.UserLoader
}

func NewUserServices(repo irepository.UserLoader) *UserServices {
	return &UserServices{repo: repo}
}

func (s *UserServices) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
func (s *UserServices) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
func (s *UserServices) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	newUserID := uuid.New()
	user.ID = newUserID
	user.Created = time.Now()
	user.Updated = time.Now()
	return s.repo.InsertUser(ctx, user)
}
