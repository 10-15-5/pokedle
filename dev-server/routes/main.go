package routes

import (
	"net/http"

	"github.com/gabr0236/pokedle/dev-server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		// curl http://localhost:3000/api/submit-guess --include --header "Content-Type: application/json" --request "GET" --data '"squirtle"'
		api.GET("/submit-guess", controllers.SubmitGuess)
		// curl http://localhost:3000/api/secret-pokemons --include --header "Content-Type: application/json" --request "GET"
		api.GET("/secret-pokemons", controllers.GetSecretPokemon)
		// curl http://localhost:3000/api/secret-pokemons --include --header "Content-Type: application/json" --request "POST"
		api.POST("/secret-pokemons", controllers.NewSecretPokemon)
	}

	// curl http://localhost:3000/health --include --header "Content-Type: application/json" --request "GET"
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running...",
		})
	})
}
