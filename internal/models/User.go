package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    Password  string    `gorm:"not null" json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword() error {
	passwordBytes := []byte(u.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	hashedPasswordBytes := []byte(u.Password)
	passwordBytes := []byte(password)
	
	return bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
}

