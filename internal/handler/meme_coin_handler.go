package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/HackHow/meme-coin/internal/common"
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

func UpdateMemeCoin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.HandleError(c, &common.AppError{
			Code:    400,
			Message: "Invalid meme coin ID",
			Data:    nil,
		})
		return
	}

	var updateReq dtos.UpdateMemeCoinRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		common.HandleError(c, &common.AppError{
			Code:    400,
			Message: "Invalid request payload",
			Data:    nil,
		})
		return
	}

	err = service.UpdateMemeCoin(uint(id), updateReq.Description)
	if err != nil {
		log.Printf("Failed to update meme coin: %v", err)
		common.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Meme coin updated successfully",
		"data":    nil,
	})
}

func DeleteMemeCoin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid meme coin ID",
			"data":    nil,
		})
		return
	}

	if err := service.DeleteMemeCoin(uint(id)); err != nil {
		log.Printf("Failed to delete meme coin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete meme coin",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Meme coin deleted successfully",
		"data":    gin.H{"id": id},
	})
}

func PokeMemeCoin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid meme coin ID",
			"data":    nil,
		})
		return
	}

	if err := service.PokeMemeCoin(uint(id)); err != nil {
		log.Printf("Failed to poke meme coin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to poke meme coin",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Meme coin poked successfully",
		"data":    nil,
	})
}
