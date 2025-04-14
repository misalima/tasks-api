package dto

import (
	"tasks-api/internal/core/domain"
)

type RegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (d RegisterDTO) ToDomain() domain.User {
	return domain.User{
		Email:    d.Email,
		Password: d.Password,
		Username: d.Username,
	}
}
