package main

import (
	"log"

	"github.com/gabr0236/pokedle/dev-server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}
	router := gin.Default()

	routes.SetupRoutes(router)

	// client = data.InitDB()
	// defer client.Disconnect(context.Background())

	router.Run("localhost:3000")
}
