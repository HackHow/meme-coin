package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/HackHow/meme-coin/config"
	"github.com/HackHow/meme-coin/internal/handler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/" + config.AppConfig.APIVersion)
	{
		api.POST("/meme-coins", handler.CreateMemeCoin)
		api.GET("/meme-coins/:id", handler.GetMemeCoin)
		api.PATCH("/meme-coins/:id", handler.UpdateMemeCoin)
		api.DELETE("/meme-coins/:id", handler.DeleteMemeCoin)
		api.POST("/meme-coins/:id/poke", handler.PokeMemeCoin)
	}

	return r
}
