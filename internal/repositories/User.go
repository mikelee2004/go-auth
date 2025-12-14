package repositories

import (
	"auth-go/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {

	if r.db == nil {
		return fmt.Errorf("database connection is nil in UserRepository")
	}

	var existingUser models.User
	result := r.db.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser)
	if result.Error == nil {
		return fmt.Errorf("this email is already taken")
	}

	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// TODO: func FindByID(id uint) (*models.User, error) {]
