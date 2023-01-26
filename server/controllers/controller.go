package controllers

import (
	"net/http"

	"github.com/gabr0236/pokedle/server/services"
	"github.com/gin-gonic/gin"
)

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
