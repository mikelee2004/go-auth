package database

import (
    "fmt"
    "log"
    
    "auth-go/internal/config"
    "auth-go/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() *gorm.DB {
	cfg := config.LoadDatabaseConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error 
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	// automigrating required models via gorm
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate models: ", err)
	}

	log.Println("Successfully connected to the database!")
	return nil
}

