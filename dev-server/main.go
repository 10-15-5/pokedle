package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gabr0236/pokedle/dev-server/database"
	"github.com/gabr0236/pokedle/dev-server/jobs"
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

	database.StartMongoDB()

	defer database.CloseMongoDB()
	// client = data.InitDB()
	// defer client.Disconnect(context.Background())

	rand.Seed(time.Now().UnixNano()) //seed rand to get diffrent values
	jobs.StartJobs()

	port := os.Getenv("SERVER_PORT")
	router.Run("localhost:" + port)
}
