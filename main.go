package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Hello! Projects Starts!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run a server: %v", err)
	}
}