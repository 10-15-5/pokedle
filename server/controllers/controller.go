package controllers

import (
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/gabr0236/pokedle/server/data"
	"github.com/gabr0236/pokedle/server/models"
	"github.com/gabr0236/pokedle/server/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func NewSecretPokemon(c *gin.Context) {
	currentSecretPokemon, _ := services.GetSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.NewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func UpdateCurrentDailyStatsWithGamesWon(c *gin.Context) {

	dailyStats, err := services.UpdateCurrentDailyStatsWithGamesWon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{"gamesWon": dailyStats.GamesWon})
}

func GetDailyStats(c *gin.Context) {

	dailyStats, err := services.GetDailyStats(c.Param("date"))

	if err != nil && err != mongo.ErrNoDocuments {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, dailyStats)

}

func GetUser(c *gin.Context) {

	userId := c.Param("userId")

	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	user := services.GetUser(mongoUserId)

	fmt.Println(user)

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func CreateUser(c *gin.Context) {

	user := models.NewUser()

	if err := services.SaveUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.SetCookie("userId", user.ID.Hex(), math.MaxInt32, "/", os.Getenv("DOMAIN"), true, false) //TODO: eventually make httpOnly=true

	c.Writer.WriteHeader(http.StatusCreated)
}
