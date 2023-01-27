package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gabr0236/pokedle/server/database"
	"github.com/gabr0236/pokedle/server/jobs"
	"github.com/gabr0236/pokedle/server/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CLIENT_URL")},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router)

	mongoClient := database.GetMongoDBClient()
	defer mongoClient.Disconnect(context.TODO())

	rand.Seed(time.Now().UnixNano()) //seed rand to get diffrent values
	jobs.StartJobs()

	port := os.Getenv("SERVER_PORT")
	router.Run("localhost:" + port)
}
