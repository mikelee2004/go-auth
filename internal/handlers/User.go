package handlers

import (
	"auth-go/internal/models"
	"auth-go/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo *repositories.UserRepository
	JWTSecret string
}

func NewAuthHandler(userRepo *repositories.UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
		JWTSecret: jwtSecret,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
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
	user := models.User{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	}

	if err := h.userRepo.CreateUser(&user); err != nil {
		if err == repositories.ErrUserExists {
			c.JSON(http.StatusConflict, gin.H{
				"error": "user with this email already exists",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create User!",
			})
		}
		return 
	}
	
	// TODO: Return success or failure response
	// If success: return user and JWT token
	// Else: return error code
}

// TODO: Login handler