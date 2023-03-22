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
		api.GET("/classic-secret-pokemons", controllers.GetClassicSecretPokemon)
		api.GET("/classic-previous-secret-pokemons", controllers.GetClassicPreviousSecretPokemon)
		api.POST("/classic-secret-pokemons", controllers.NewClassicSecretPokemon)

		api.GET("/flavortext-secret-pokemons", controllers.GetFlavortextSecretPokemon)
		api.GET("/flavortext-previous-secret-pokemons", controllers.GetFlavortextPreviousSecretPokemon)
		api.POST("/flavortext-secret-pokemons", controllers.NewFlavortextSecretPokemon)

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
			"message": "Server v0.4 DEV is running...",
		})
	})
}
