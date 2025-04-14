package iservice

import "tasks-api/internal/core/domain"

type AuthManager interface {
	Login(email, password string) (string, error)
	Logout(token string) error
	Register(user domain.User) (*domain.User, error)
}
