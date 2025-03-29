package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/HackHow/meme-coin/internal/handler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/meme-coins", handler.CreateMemeCoin)
	r.GET("/meme-coins/:id", handler.GetMemeCoin)

	return r
}
