// cmd/server/main.go
package main

import (
	"log"

	"github.com/HackHow/meme-coin/config"
	"github.com/HackHow/meme-coin/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "Welcome to Meme Coin API!",
			"data":    nil,
		})
	})

	port := config.AppConfig.ServerPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running at http://localhost:%s/", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
