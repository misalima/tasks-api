package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/iservice"
	"tasks-api/internal/http/dto"
)

type AuthHandler struct {
	authManager iservice.AuthManager
}

func NewAuthHandler(authManager iservice.AuthManager) *AuthHandler {
	return &AuthHandler{
		authManager: authManager,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request data"})
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request data"})
	}

	token, err := h.authManager.Login(req.Email, req.Password)
	if err != nil {
		if err.Error() == "invalid password" || errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
	}
	return c.JSON(http.StatusOK, dto.LoginResponseDTO{Token: token})
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request data"})
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request data"})
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid email format"})
	}

	user, err := h.authManager.Register(req.ToDomain())
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return c.JSON(http.StatusConflict, dto.ErrorResponse{Message: "Email already taken"})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
	}
	return c.JSON(http.StatusCreated, dto.UserFromDomain(*user))
}
