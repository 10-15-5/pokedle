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

// CLASSIC
func GetClassicSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetClassicSecretPokemon()

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

func NewClassicSecretPokemon(c *gin.Context) {

	if os.Getenv("ENVIRONMENT") == "production" {
		c.JSON(http.StatusBadRequest, "This operation is not supported in production environment")
		return
	}

	currentSecretPokemon, _ := services.GetClassicSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.ClassicNewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)

}

// FLAVORTEXT
func GetFlavortextSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetFlavortextSecretPokemon()

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

func NewFlavortextSecretPokemon(c *gin.Context) {

	if os.Getenv("ENVIRONMENT") == "production" {
		c.JSON(http.StatusBadRequest, "This operation is not supported in production environment")
		return
	}

	currentSecretPokemon, _ := services.GetClassicSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.FlavortextNewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)

}

// SILHOUETTE
func GetSilhouetteSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetSilhouetteSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func GetSilhouettePreviousSecretPokemon(c *gin.Context) {
	secretPokemon, err := services.GetSilhouettePreviousSecretPokemon()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, secretPokemon)
}

func NewSilhouetteSecretPokemon(c *gin.Context) {

	if os.Getenv("ENVIRONMENT") == "production" {
		c.JSON(http.StatusBadRequest, "This operation is not supported in production environment")
		return
	}

	currentSecretPokemon, _ := services.GetSilhouetteSecretPokemon()

	pokemonData := data.GetPokemonData()

	secretPokemon := services.SilhouetteNewSecretPokemon(currentSecretPokemon, pokemonData)

	c.IndentedJSON(http.StatusOK, secretPokemon)

}

type NumberOfGuessesRequest struct {
	NumberOfGuesses int `json:"numberOfGuesses" bson:"numberOfGuesses" binding:"required"`
}

func UpdateClassicCurrentDailyStatsGamesWon(c *gin.Context) {

	var updateUserGameWonRequest NumberOfGuessesRequest

	if err := c.ShouldBindJSON(&updateUserGameWonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dailyStats, err := services.UpdateClassicCurrentDailyStatsGamesWon(updateUserGameWonRequest.NumberOfGuesses)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"dailyStats": dailyStats})
}

func UpdateFlavortextCurrentDailyStatsGamesWon(c *gin.Context) {

	var updateUserGameWonRequest NumberOfGuessesRequest

	if err := c.ShouldBindJSON(&updateUserGameWonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dailyStats, err := services.UpdateFlavortextCurrentDailyStatsGamesWon(updateUserGameWonRequest.NumberOfGuesses)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"dailyStats": dailyStats})
}

func UpdateSilhouetteCurrentDailyStatsGamesWon(c *gin.Context) {

	var updateUserGameWonRequest NumberOfGuessesRequest

	if err := c.ShouldBindJSON(&updateUserGameWonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dailyStats, err := services.UpdateSilhouetteCurrentDailyStatsGamesWon(updateUserGameWonRequest.NumberOfGuesses)

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

func UpdateUserStreaks(c *gin.Context) {

	userId := c.Param("userId")

	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := services.GetUser(mongoUserId)

	var streakClassic = services.CalculateStreak(user.ClassicGamesWon, user.ClassicCurrentStreak)
	var streakFlavortext = services.CalculateStreak(user.FlavortextGamesWon, user.FlavortextCurrentStreak)
	var streakSilhouette = services.CalculateStreak(user.SilhouetteGamesWon, user.SilhouetteCurrentStreak)

	if streakClassic != user.ClassicCurrentStreak || streakFlavortext != user.FlavortextCurrentStreak || streakSilhouette != user.SilhouetteCurrentStreak {
		services.UpdateUserStreaks(mongoUserId, streakClassic, streakFlavortext, streakSilhouette)
		c.AbortWithStatusJSON(200, gin.H{"msg": "Streak Updated"})
		return
	}

	c.AbortWithStatus(200)
}

func HandleUserClassicGameWon(c *gin.Context) {

	userId := c.Param("userId")
	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, "UserId string is not a valid ObjectID")
		return
	}

	user := services.GetUser(mongoUserId)

	UpdateUserGameWon(c, services.Classic, user.ClassicGamesWon, user.ID, user.ClassicCurrentStreak, user.ClassicMaxStreak)
}

func HandleUserFlavortextGameWon(c *gin.Context) {

	userId := c.Param("userId")
	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, "UserId string is not a valid ObjectID")
		return
	}

	user := services.GetUser(mongoUserId)

	UpdateUserGameWon(c, services.Flavortext, user.FlavortextGamesWon, user.ID, user.FlavortextCurrentStreak, user.FlavortextMaxStreak)
}

func HandleUserSilhouetteGameWon(c *gin.Context) {

	userId := c.Param("userId")
	mongoUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, "UserId string is not a valid ObjectID")
		return
	}

	user := services.GetUser(mongoUserId)

	UpdateUserGameWon(c, services.Silhouette, user.SilhouetteGamesWon, user.ID, user.SilhouetteCurrentStreak, user.SilhouetteMaxStreak)
}

func UpdateUserGameWon(c *gin.Context, gameMode string, gamesWon []models.GameWon, userId primitive.ObjectID, currentStreak int, maxStreak int) {

	var updateUserGameWonRequest NumberOfGuessesRequest

	if err := c.ShouldBindJSON(&updateUserGameWonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateUserGameWonRequest.NumberOfGuesses <= 0 {
		c.JSON(http.StatusBadRequest, "Number of guesses cannot be <= 0")
		return
	}

	gameWon := models.GameWon{
		NumberOfGuesses: updateUserGameWonRequest.NumberOfGuesses,
		CreatedAt:       time.Now(),
	}

	var streak int

	if len(gamesWon) > 0 {

		lastGameWon := gamesWon[len(gamesWon)-1]

		if services.DateEqual(time.Now(), lastGameWon.CreatedAt) {
			c.JSON(http.StatusBadRequest, "User has already won a game today")
			return
		}

		streak = services.CalculateStreak(gamesWon, currentStreak)
	}

	streak++

	newMaxStreak := int(math.Max(float64(currentStreak), float64(maxStreak)))

	isFirstTryWin := services.If(updateUserGameWonRequest.NumberOfGuesses == 1, 1, 0)

	updatedUser := services.FindAndUpdateUserWithGameWon(userId, gameWon, streak, newMaxStreak, isFirstTryWin, gameMode)

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}
