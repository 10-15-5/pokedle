package main

import (
	"net/http"

	"github.com/gabr0236/pokedle/dev-server/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	SetupMongoDB()

	//curl http://localhost:3000/api/submit-guess --include --header "Content-Type: application/json" --request "GET" --data '"squirtle"'
	api := router.Group("/api")
	{
		api.GET("/submit-guess", submitGuess)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run("localhost:3000")
}

/*
curl http://localhost:3000/guess --include --header "Content-Type: application/json" --request "GET" --data '"squirtle"'
*/

// getAlbums responds with the list of all albums as JSON.
func submitGuess(c *gin.Context) {
	var name string
	println("Call enter")

	if err := c.BindJSON(&name); err != nil {
		return
	}

	services.SubmitGuess(name)

	c.IndentedJSON(http.StatusOK, name)
}
