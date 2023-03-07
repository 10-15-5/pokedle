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
		api.GET("/previous-secret-pokemons", controllers.GetPreviousSecretPokemon)
		// curl http://localhost:3000/api/secret-pokemons --include --header "Content-Type: application/json" --request "POST"
		api.POST("/secret-pokemons", controllers.NewSecretPokemon)

		// curl http://localhost:3000/api/games-won --include --header "Content-Type: application/json" --request "POST"
		api.POST("/games-won", controllers.UpdateCurrentDailyStatsWithGamesWon)

		// curl http://localhost:3000/api/daily-stats/21321 --request "GET"
		api.GET("/daily-stats/:date", controllers.GetDailyStats)

		api.GET("/users/:userId", controllers.GetUser)

		api.PATCH("/users/:userId", controllers.UpdateUserStreak)

		api.POST("/users", controllers.CreateUser)

		api.POST("/users/:userId/gamesWon", controllers.UpdateUserGameWon)
	}

	// curl http://localhost:3000/health --include --header "Content-Type: application/json" --request "GET"
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server v0.3 is running...",
		})
	})
}
