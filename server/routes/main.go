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
		api.POST("/classic-games-won", controllers.UpdateClassicCurrentDailyStatsGamesWon)
		api.POST("/users/:userId/classic-wins", controllers.HandleUserClassicGameWon)

		api.GET("/flavortext-secret-pokemons", controllers.GetFlavortextSecretPokemon)
		api.GET("/flavortext-previous-secret-pokemons", controllers.GetFlavortextPreviousSecretPokemon)
		api.POST("/flavortext-secret-pokemons", controllers.NewFlavortextSecretPokemon)
		api.POST("/flavortext-games-won", controllers.UpdateFlavortextCurrentDailyStatsGamesWon)
		api.POST("/users/:userId/flavortext-wins", controllers.HandleUserFlavortextGameWon)

		api.GET("/silhouette-secret-pokemons", controllers.GetSilhouetteSecretPokemon)
		api.GET("/silhouette-previous-secret-pokemons", controllers.GetSilhouettePreviousSecretPokemon)
		api.POST("/silhouette-secret-pokemons", controllers.NewSilhouetteSecretPokemon)
		api.POST("/silhouette-games-won", controllers.UpdateSilhouetteCurrentDailyStatsGamesWon)
		api.POST("/users/:userId/silhouette-wins", controllers.HandleUserSilhouetteGameWon)

		api.GET("/daily-stats/:date", controllers.GetDailyStats)

		api.GET("/users/:userId", controllers.GetUser)
		api.PATCH("/users/:userId", controllers.UpdateUserStreaks) //To update streak on load
		api.POST("/users", controllers.CreateUser)

	}

	// curl http://localhost:3000/health --include --header "Content-Type: application/json" --request "GET"
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pokedle server v2.0 running...",
		})
	})
}
