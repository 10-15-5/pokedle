package controllers

import (
	"net/http"

	"github.com/gabr0236/pokedle/server/data"
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
	currentSecretPokemon, _ := services.GetSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.NewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)
}
