package controllers

import (
	"net/http"

	"github.com/gabr0236/pokedle/dev-server/services"
	"github.com/gin-gonic/gin"
)

func SubmitGuess(c *gin.Context) {
	// var name string
	// println("Call enter")

	// if err := c.BindJSON(&name); err != nil {
	// 	return
	// }

	// services.SubmitGuess(name)

	// c.IndentedJSON(http.StatusOK, name)
}

func GetSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetSecretPokemon()

	if err != nil {
		c.AbortWithStatus(500)
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func NewSecretPokemon(c *gin.Context) {
	err := services.NewSecretPokemon()

	if err != nil {
		c.AbortWithStatus(500)
	}

	c.Status(200)
}
