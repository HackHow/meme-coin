package handler

import (
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateMemeCoin(c *gin.Context) {
	var coin model.MemeCoin
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request payload",
			"data":    nil,
		})
		return
	}

	if err := service.CreateMemeCoin(&coin); err != nil {
		log.Printf("Failed to create meme coin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create meme coin",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "Meme coin created successfully",
		"data":    coin,
	})
}
