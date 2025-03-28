package handler

import (
	"fmt"

	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/HackHow/meme-coin/internal/dtos"
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/internal/service"
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

func GetMemeCoin(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Print("idStr:", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid meme coin ID",
			"data":    nil,
		})
		return
	}

	coin, err := service.GetMemeCoin(uint(id))
	fmt.Print("coin:", coin)
	fmt.Print("err:", err)
	if err != nil {
		log.Printf("Failed to get meme coin: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Meme coin not found",
			"data":    nil,
		})
		return
	}

	response := dtos.GetMemeCoinResponse{
		Name:            coin.Name,
		Description:     coin.Description,
		CreatedAt:       coin.CreatedAt,
		PopularityScore: coin.PopularityScore,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Meme coin details fetched successfully",
		"data":    response,
	})
}
