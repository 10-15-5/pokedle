package controllers

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

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
	//TODO: increment daily first try wins
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

type UpdateUserGameWonRequest struct {
	NumberOfGuesses int `json:"numberOfGuesses"`
}

func UpdateUserGameWon(c *gin.Context) {

	userId := c.Param("userId")

	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, "UserId string is not a valid ObjectID")
		return
	}

	user := services.GetUser(mongoUserId)

	if len(user.GamesWon) > 0 {

		lastGameWon := user.GamesWon[len(user.GamesWon)-1]

		//TODO: test this and then comment out untill 1 guess a day is implemented
		if services.DateEqual(time.Now(), lastGameWon.CreatedAt) {
			c.JSON(http.StatusBadRequest, "User has already won a game today")
			return
		}

	}

	var updateUserGameWonRequest UpdateUserGameWonRequest

	c.BindJSON(&updateUserGameWonRequest)

	if updateUserGameWonRequest.NumberOfGuesses <= 0 {
		c.JSON(http.StatusBadRequest, "Number of guesses cannot be <= 0")
		return
	}

	gameWon := models.GameWon{
		NumberOfGuesses: updateUserGameWonRequest.NumberOfGuesses,
		CreatedAt:       time.Now(),
	}

	streak := services.CalculateStreak(user.GamesWon, user.CurrentStreak)

	maxStreak := int(math.Max(float64(streak), float64(user.MaxStreak)))

	isFirstTryWin := services.If(updateUserGameWonRequest.NumberOfGuesses == 1, 1, 0)

	updatedUser := services.FindAndUpdateUserWithGameWon(mongoUserId, gameWon, streak, maxStreak, isFirstTryWin)

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}
