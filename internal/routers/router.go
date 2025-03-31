package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/HackHow/meme-coin/config"
	"github.com/HackHow/meme-coin/internal/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("/api/" + config.AppConfig.APIVersion)
	{
		api.POST("/meme-coins", handlers.CreateMemeCoin)
		api.GET("/meme-coins/:id", handlers.GetMemeCoin)
		api.PATCH("/meme-coins/:id", handlers.UpdateMemeCoin)
		api.DELETE("/meme-coins/:id", handlers.DeleteMemeCoin)
		api.POST("/meme-coins/:id/poke", handlers.PokeMemeCoin)
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	return r
}
