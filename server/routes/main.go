package routes

import (
	"net/http"

	"github.com/gabr0236/pokedle/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//TODO: split routes
	api := router.Group("/api")
	{
		// curl http://localhost:3000/api/secret-pokemons --include --header "Content-Type: application/json" --request "GET"
		api.GET("/secret-pokemons", controllers.GetSecretPokemon)
		// curl http://localhost:3000/api/secret-pokemons --include --header "Content-Type: application/json" --request "POST"
		api.POST("/secret-pokemons", controllers.NewSecretPokemon)

		// curl http://localhost:3000/api/games-won --include --header "Content-Type: application/json" --request "POST"
		api.POST("/games-won", controllers.UpdateCurrentDailyStatsWithGamesWon)

		// curl http://localhost:3000/api/daily-stats/21321 --request "GET"
		api.GET("/daily-stats/:date", controllers.GetDailyStats)
	}

	// curl http://localhost:3000/health --include --header "Content-Type: application/json" --request "GET"
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running...",
		})
	})
}
