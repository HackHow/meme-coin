// cmd/server/main.go
package main

import (
	"log"

	"github.com/HackHow/meme-coin/config"
	"github.com/HackHow/meme-coin/internal/routers"
	"github.com/HackHow/meme-coin/pkg/database"
)

func main() {
	config.LoadConfig()
	database.InitDB()

	router := routers.InitRouter()

	port := config.AppConfig.ServerPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running at http://localhost:%s/", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
