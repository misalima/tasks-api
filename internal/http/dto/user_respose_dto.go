package dto

import (
	"github.com/google/uuid"
	"tasks-api/internal/core/domain"
	"time"
)

type UserResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func UserFromDomain(user domain.User) UserResponseDTO {
	return UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Created:  user.Created,
		Updated:  user.Updated,
	}
}
