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

func GetClassicSecretPokemon() (models.Pokemon, error) {

	r := secretPokemon.GetPokemonRepository(MongoClient)

	return r.FindCurrentSecretPokemon(context.TODO(), ClassicSecretPokemons)
}

func GetFlavortextSecretPokemon() (models.Pokemon, error) {

	r := secretPokemon.GetPokemonRepository(MongoClient)

	return r.FindCurrentSecretPokemon(context.TODO(), FlavortextSecretPokemons)
}

func GetClassicPreviousSecretPokemon() (models.Pokemon, error) {

	r := secretPokemon.GetPokemonRepository(MongoClient)

	secretPokemons, err := r.FindLastTwoSecretPokemon(context.TODO(), ClassicSecretPokemons)

	return secretPokemons[1], err
}

func GetFlavortextPreviousSecretPokemon() (models.Pokemon, error) {

	r := secretPokemon.GetPokemonRepository(MongoClient)

	secretPokemons, err := r.FindLastTwoSecretPokemon(context.TODO(), FlavortextSecretPokemons)

	return secretPokemons[1], err
}

func ClassicNewSecretPokemon(currentSecretPokemon models.Pokemon, pokemonData []models.Pokemon) models.Pokemon {
	return newSecretPokemon(currentSecretPokemon, pokemonData, ClassicSecretPokemons)
}

func FlavortextNewSecretPokemon(currentSecretPokemon models.Pokemon, pokemonData []models.Pokemon) models.Pokemon {
	return newSecretPokemon(currentSecretPokemon, pokemonData, FlavortextSecretPokemons)
}

func newSecretPokemon(currentSecretPokemon models.Pokemon, pokemonData []models.Pokemon, collection string) models.Pokemon {
	containsRandomPokemon := true

	var randomPokemon models.Pokemon

	for containsRandomPokemon {
		randomPokemon = pokemonData[rand.Intn(len(pokemonData))]
		fmt.Println(randomPokemon)

		containsRandomPokemon = randomPokemon.Name == currentSecretPokemon.Name
	}

	fmt.Println(randomPokemon)

	r := secretPokemon.GetPokemonRepository(MongoClient)

	r.InsertNewPokemon(context.TODO(), randomPokemon, collection)

	fmt.Println("SecretPokemon Updated.")
	return randomPokemon
}

func UpdateClassicDailySecretPokemon() error {
	return UpdateDailySecretPokemon(ClassicSecretPokemons)
}
func UpdateFlavortextDailySecretPokemon() error {
	return UpdateDailySecretPokemon(FlavortextSecretPokemons)
}

func UpdateDailySecretPokemon(collection string) error {
	pokemonData := data.GetPokemonData()

	if pokemonData == nil {
		panic("PokemonData not found")
	}

	r := secretPokemon.GetPokemonRepository(MongoClient)

	recentSecretPokemons := r.FindLastSecretPokemons(context.TODO(), 30, collection)

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

	_, err := r.InsertNewPokemon(context.TODO(), randomPokemon, ClassicSecretPokemons)
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

func CalculateStreak(lastGameWon models.ClassicGameWon, currStreak int) int {

	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)

	if DateEqual(yesterday, lastGameWon.CreatedAt) || DateEqual(today, lastGameWon.CreatedAt) {
		return currStreak
	}

	return 0
}

func FindAndUpdateUserWithGameWon(
	userId primitive.ObjectID,
	gameWon models.ClassicGameWon,
	currentStreak int,
	maxStreak int,
	isFirstTryWin int,
) models.User {

	r := user.GetUserRepository(MongoClient)

	return r.InsertNewGameWon(userId, gameWon, currentStreak, maxStreak, isFirstTryWin)
}

func UpdateUserStreak(
	userId primitive.ObjectID,
	currentStreak int,
) {

	r := user.GetUserRepository(MongoClient)

	r.UpdateUserStreak(userId, currentStreak)
}
