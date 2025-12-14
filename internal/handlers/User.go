package handlers

import (
	"auth-go/internal/models"
	"auth-go/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo *repositories.UserRepository
}

func NewAuthHandler(userRepo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// validating register data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Failed to hash password!",
		})
		return
	}

	if err := h.userRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// TODO: Login handler
