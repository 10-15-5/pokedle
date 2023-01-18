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
	}

	// curl http://localhost:3000/ping --include --header "Content-Type: application/json" --request "GET"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
