package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gabr0236/pokedle/server/data"
	"github.com/gabr0236/pokedle/server/jobs"
	"github.com/gabr0236/pokedle/server/routes"
	"github.com/gabr0236/pokedle/server/services"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//TODO: this will not work with docker, although somehow env's are still set? 🤨
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}

	log.Println("ENV TEST: GET SERVER PORT: " + os.Getenv("SERVER_PORT"))

	log.Println("ENV TEST: GET CLIENT_URL: " + os.Getenv("CLIENT_URL"))

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CLIENT_URL")},
		AllowMethods:     []string{"GET", "POST", "UPDATE", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router)

	mongoClient := data.GetMongoDBClient()
	services.MongoClient = mongoClient
	defer mongoClient.Disconnect(context.TODO())

	rand.Seed(time.Now().UnixNano()) //seed rand to get diffrent values
	jobs.StartJobs()

	port := os.Getenv("SERVER_PORT")
	router.Run(":" + port)
}
