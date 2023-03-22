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

func GetClassicSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetClassicSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func GetFlavortextSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetFlavortextSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func GetClassicPreviousSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetClassicPreviousSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func GetFlavortextPreviousSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetFlavortextPreviousSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func NewClassicSecretPokemon(c *gin.Context) {

	if os.Getenv("ENVIRONMENT") == "production" {
		c.JSON(http.StatusBadRequest, "This operation is not supported in production environment")
		return
	}

	currentSecretPokemon, _ := services.GetClassicSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.ClassicNewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)

	//TODO: also reset dailySecretPokemon for Flavortext
}

func NewFlavortextSecretPokemon(c *gin.Context) {

	if os.Getenv("ENVIRONMENT") == "production" {
		c.JSON(http.StatusBadRequest, "This operation is not supported in production environment")
		return
	}

	currentSecretPokemon, _ := services.GetClassicSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.FlavortextNewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)

	//TODO: also reset dailySecretPokemon for Flavortext
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"dailyStats": dailyStats})
}

func GetDailyStats(c *gin.Context) {

	dailyStats, err := services.GetDailyStats(c.Param("date"))

	if err != nil && err != mongo.ErrNoDocuments {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dailyStats)

}

func CreateUser(c *gin.Context) {

	user := models.NewUser()

	if err := services.SaveUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
		return
	}

	user := services.GetUser(mongoUserId)

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func UpdateUserStreak(c *gin.Context) {

	userId := c.Param("userId")

	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := services.GetUser(mongoUserId)

	var streak int

	if len(user.ClassicGamesWon) > 0 {

		lastGameWon := user.ClassicGamesWon[len(user.ClassicGamesWon)-1]

		streak = services.CalculateStreak(lastGameWon, user.ClassicCurrentStreak)
	}

	if streak != user.ClassicCurrentStreak {
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

	gameWon := models.ClassicGameWon{
		NumberOfGuesses: updateUserGameWonRequest.NumberOfGuesses,
		CreatedAt:       time.Now(),
	}

	var streak int

	if len(user.ClassicGamesWon) > 0 {

		lastGameWon := user.ClassicGamesWon[len(user.ClassicGamesWon)-1]

		//TODO: test this and then comment out untill 1 guess a day is implemented
		if services.DateEqual(time.Now(), lastGameWon.CreatedAt) {
			c.JSON(http.StatusBadRequest, "User has already won a game today")
			return
		}

		streak = services.CalculateStreak(lastGameWon, user.ClassicCurrentStreak)
	}

	streak++

	maxStreak := int(math.Max(float64(streak), float64(user.ClassicMaxStreak)))

	isFirstTryWin := services.If(updateUserGameWonRequest.NumberOfGuesses == 1, 1, 0)

	updatedUser := services.FindAndUpdateUserWithGameWon(mongoUserId, gameWon, streak, maxStreak, isFirstTryWin)

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}
