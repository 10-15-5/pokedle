package controllers

import (
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

type NumberOfGuessesRequest struct {
	NumberOfGuesses int `json:"numberOfGuesses" bson:"numberOfGuesses" binding:"required"`
}

func GetSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func GetPreviousSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetPreviousSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func NewSecretPokemon(c *gin.Context) {

	if os.Getenv("ENVIRONMENT") == "production" {
		c.JSON(http.StatusBadRequest, "This operation is not supported in production environment")
		return
	}

	currentSecretPokemon, _ := services.GetSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.NewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func UpdateCurrentDailyStatsWithGamesWon(c *gin.Context) {

	var updateUserGameWonRequest NumberOfGuessesRequest

	if err := c.ShouldBindJSON(&updateUserGameWonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dailyStats, err := services.UpdateCurrentDailyStatsWithGamesWon(updateUserGameWonRequest.NumberOfGuesses)
	//TODO: increment daily first try wins
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{"dailyStats": dailyStats})
}

func GetDailyStats(c *gin.Context) {

	dailyStats, err := services.GetDailyStats(c.Param("date"))

	if err != nil && err != mongo.ErrNoDocuments {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, dailyStats)

}

func CreateUser(c *gin.Context) {

	user := models.NewUser()

	if err := services.SaveUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie("userId", user.ID.Hex(), math.MaxInt32, "/", os.Getenv("DOMAIN"), true, false)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func GetUser(c *gin.Context) {

	userId := c.Param("userId")

	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	user := services.GetUser(mongoUserId)

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func UpdateUserStreak(c *gin.Context) {

	userId := c.Param("userId")

	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	user := services.GetUser(mongoUserId)

	var streak int

	if len(user.GamesWon) > 0 {

		lastGameWon := user.GamesWon[len(user.GamesWon)-1]

		streak = services.CalculateStreak(lastGameWon, user.CurrentStreak)
	}

	if streak != user.CurrentStreak {
		services.UpdateUserStreak(mongoUserId, streak)
		c.AbortWithStatusJSON(200, gin.H{"msg": "Streak Updated"})
	}

	c.AbortWithStatus(200)
}

func UpdateUserGameWon(c *gin.Context) {

	var updateUserGameWonRequest NumberOfGuessesRequest

	if err := c.ShouldBindJSON(&updateUserGameWonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateUserGameWonRequest.NumberOfGuesses <= 0 {
		c.JSON(http.StatusBadRequest, "Number of guesses cannot be <= 0")
		return
	}

	userId := c.Param("userId")
	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, "UserId string is not a valid ObjectID")
		return
	}

	user := services.GetUser(mongoUserId)

	gameWon := models.GameWon{
		NumberOfGuesses: updateUserGameWonRequest.NumberOfGuesses,
		CreatedAt:       time.Now(),
	}

	var streak int

	if len(user.GamesWon) > 0 {

		lastGameWon := user.GamesWon[len(user.GamesWon)-1]

		//TODO: test this and then comment out untill 1 guess a day is implemented
		if services.DateEqual(time.Now(), lastGameWon.CreatedAt) {
			c.JSON(http.StatusBadRequest, "User has already won a game today")
			return
		}

		streak = services.CalculateStreak(lastGameWon, user.CurrentStreak)
	}

	streak++

	maxStreak := int(math.Max(float64(streak), float64(user.MaxStreak)))

	isFirstTryWin := services.If(updateUserGameWonRequest.NumberOfGuesses == 1, 1, 0)

	updatedUser := services.FindAndUpdateUserWithGameWon(mongoUserId, gameWon, streak, maxStreak, isFirstTryWin)

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}
