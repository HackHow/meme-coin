package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "Welcome to Meme Coin API!",
			"data":    nil,
		})
	})

	if err := router.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
