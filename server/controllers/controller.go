package controllers

import (
	"fmt"
	"net/http"

	"github.com/gabr0236/pokedle/server/data"
	"github.com/gabr0236/pokedle/server/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

// Dont allow for param/payload here for security reasons
func UpdateCurrentDailyStatsWithGamesWon(c *gin.Context) {
	//TODO: add guard so single user cant spam this, use user cookie to identify, maybe even IP identification
	dailyStats, err := services.UpdateCurrentDailyStatsWithGamesWon()

	if err != nil {
		c.AbortWithStatus(500)
	}

	c.IndentedJSON(http.StatusOK, dailyStats.GamesWon)
}

func GetDailyStats(c *gin.Context) {

	dailyStats, err := services.GetDailyStats(c.Param("date"))

	fmt.Println(err)

	if err != nil && err != mongo.ErrNoDocuments {
		c.AbortWithStatus(500)
		return
	}

	c.IndentedJSON(http.StatusOK, dailyStats)

}
