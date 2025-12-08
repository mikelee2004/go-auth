package main

import (
	"auth-go/app"
	"auth-go/internal/database"
	"auth-go/internal/handlers"
	"auth-go/internal/repositories"
	"log"
)

func main() {
    // Инициализация базы данных
    db := database.InitializeDB()

    // Инициализация репозитория
    userRepo := repositories.NewUserRepository(db)
    
    // Инициализация хендлера
    authHandler := handlers.NewAuthHandler(userRepo)
    
    // Настройка роутера
    router := app.SetupRouter(*authHandler)

    log.Println("Server starting on :8080")
    router.Run(":8080")
}