package main

import (
	"auth-go/internal/database"
	"auth-go/internal/handlers"
	"auth-go/internal/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitializeDB()

	r := gin.Default()

	// repositories
	userRepository := repositories.NewUserRepository(db)

	// handlers
	authHandler := handlers.NewAuthHandler(userRepository)

	// routes (TODO: Move them to the app/routes.go dir)
	r.POST("/register", authHandler.Register)

	// server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: ", err)
	}
}