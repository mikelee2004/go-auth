package app

import (
	"auth-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler handlers.AuthHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		public := api.Group("/auth")
		{
			public.POST("/register", userHandler.Register)
		}
	}
	return router
}