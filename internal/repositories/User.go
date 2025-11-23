package repositories

import (
	"auth-go/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists = errors.New("user already exists")
)

func (r *UserRepository) CreateUser(user *models.User) error {
	result := r.DB.Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrUserExists
		}
		return result.Error
	}
	return nil
}

// TODO: func FindByID(id uint) (*models.User, error) {]