package main

import (
	"auth-go/internal/config"
	"auth-go/internal/database"
	"auth-go/internal/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadDatabaseConfig()
	db := database.InitializeDB(cfg)
	
	userRepo := repositories.UserRepository(db)
	r := gin.Default()
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run a server: %v", err)
	}
}