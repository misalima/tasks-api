package services

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/iservice"
	"tasks-api/internal/http/auth"
	"time"
)

var _ iservice.AuthManager = (*AuthServices)(nil)

type AuthServices struct {
	userServices *UserServices
}

func NewAuthService(userServices *UserServices) *AuthServices {
	return &AuthServices{
		userServices: userServices,
	}
}

func (s *AuthServices) Login(email, password string) (string, error) {
	user, err := s.userServices.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	return auth.GenerateToken(user.ID, user.Username)
}
func (s *AuthServices) Logout(token string) error {
	panic("Not Implemented!")
}
func (s *AuthServices) Register(user domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	existingUser, err := s.userServices.GetUserByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return s.userServices.CreateUser(context.Background(), user)
}
