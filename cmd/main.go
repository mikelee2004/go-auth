package main

import (
	"auth-go/app"
	"auth-go/internal/database"
	"auth-go/internal/handlers"
	"auth-go/internal/repositories"
	"log"
)

func main() {
    // db
    db := database.InitializeDB()

    // repositories
    userRepo := repositories.NewUserRepository(db)
    
    // handels
    authHandler := handlers.NewAuthHandler(userRepo)
    
    // router
    router := app.SetupRouter(*authHandler)

    log.Println("Server starting on :8080")
    router.Run(":8080")
}