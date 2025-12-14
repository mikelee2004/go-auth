package main

import (
	"auth-go/api"
	"auth-go/internal/database"
	"auth-go/internal/handlers"
	"auth-go/internal/repositories"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("couldn't load env file: ", err)
	}
	// db
	db := database.InitializeDB()

	// repositories
	userRepo := repositories.NewUserRepository(db)

	// handels
	authHandler := handlers.NewAuthHandler(userRepo)

	// router
	router := api.SetupRouter(*authHandler)

	log.Println("Server starting on :8080")
	router.Run(":8080")
}
