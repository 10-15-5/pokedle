package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gabr0236/pokedle/server/data"
	"github.com/gabr0236/pokedle/server/models"
	"github.com/gabr0236/pokedle/server/repository/dailyStats"
	"github.com/gabr0236/pokedle/server/repository/secretPokemon"
	"github.com/gabr0236/pokedle/server/repository/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

func GetSecretPokemon() (models.Pokemon, error) {

	r := secretPokemon.GetPokemonRepository(MongoClient)

	return r.FindCurrentSecretPokemon(context.TODO())
}

func NewSecretPokemon(currentSecretPokemon models.Pokemon, pokemonData []models.Pokemon) models.Pokemon {
	containsRandomPokemon := true

	var randomPokemon models.Pokemon

	for containsRandomPokemon {
		randomPokemon = pokemonData[rand.Intn(len(pokemonData))]
		fmt.Println(randomPokemon)

		containsRandomPokemon = randomPokemon.Name == currentSecretPokemon.Name
	}

	r := secretPokemon.GetPokemonRepository(MongoClient)

	r.InsertNewPokemon(context.TODO(), randomPokemon)

	fmt.Println("SecretPokemon Updated.")
	return randomPokemon
}

func UpdateDailySecretPokemon() error {
	pokemonData := data.GetPokemonData()

	if pokemonData == nil {
		panic("PokemonData not found")
	}

	r := secretPokemon.GetPokemonRepository(MongoClient)

	recentSecretPokemons := r.FindRecentSecretPokemons(context.TODO(), 30)

	//TODO: if recentSecretPokemons.length > 30 remove oldest doc

	//TODO: move to helper package
	containsPokemonName := func(name string, a []models.Pokemon) bool {
		for _, p := range recentSecretPokemons {
			if p.Name == name {
				return true
			}
		}
		return false
	}

	containsRandomPokemon := true

	var randomPokemon models.Pokemon

	for containsRandomPokemon {
		randomPokemon = pokemonData[rand.Intn(len(pokemonData))]
		fmt.Println(randomPokemon)

		containsRandomPokemon = containsPokemonName(randomPokemon.Name, recentSecretPokemons)
	}

	_, err := r.InsertNewPokemon(context.TODO(), randomPokemon)
	if err != nil {
		fmt.Println("Error when inserting pokemon")
	}

	fmt.Println("SecretPokemon Updated.")

	return err
}

func UpdateCurrentDailyStatsWithGamesWon(numberOfGuesses int) (models.DailyStats, error) {
	r := dailyStats.GetPokemonRepository(MongoClient)

	currentDate := time.Now().Format("2006-01-02")

	isFirstTryWin := If(numberOfGuesses == 1, 1, 0)

	result, err := r.UpdateDailyGuessCount(context.TODO(), currentDate, isFirstTryWin, numberOfGuesses)

	return result, err
}

func GetDailyStats(date string) (models.DailyStats, error) {
	r := dailyStats.GetPokemonRepository(MongoClient)

	dailyStats, err := r.GetDailyStats(context.TODO(), date)

	return dailyStats, err
}

func GetUser(userId primitive.ObjectID) models.User {
	r := user.GetUserRepository(MongoClient)

	return r.GetUser(userId)
}

func SaveUser(newUser models.User) error {
	r := user.GetUserRepository(MongoClient)

	return r.InsertNewUser(newUser)
}

func CalculateStreak(lastGameWon models.GameWon, currStreak int) int {

	yesterday := time.Now().Add(-24 * time.Hour)

	if DateEqual(yesterday, lastGameWon.CreatedAt) {
		return currStreak + 1
	}

	return 1
}

func FindAndUpdateUserWithGameWon(
	userId primitive.ObjectID,
	gameWon models.GameWon,
	currentStreak int,
	maxStreak int,
	isFirstTryWin int,
) models.User {

	r := user.GetUserRepository(MongoClient)

	return r.InsertNewGameWon(userId, gameWon, currentStreak, maxStreak, isFirstTryWin)
}
