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
)

func GetSecretPokemon() (models.Pokemon, error) {

	mongoClient := data.GetMongoDBClient()

	r := secretPokemon.GetPokemonRepository(mongoClient)

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

	mongoClient := data.GetMongoDBClient()

	r := secretPokemon.GetPokemonRepository(mongoClient)

	r.InsertNewPokemon(context.TODO(), randomPokemon)

	fmt.Println("SecretPokemon Updated.")
	return randomPokemon
}

func UpdateDailySecretPokemon() error {
	pokemonData := data.GetPokemonData()

	mongoClient := data.GetMongoDBClient()

	if pokemonData == nil || mongoClient == nil {
		panic("Missing pokemonData or mongoClient")
	}

	r := secretPokemon.GetPokemonRepository(mongoClient)

	recentSecretPokemons, err := r.FindRecentSecretPokemons(context.TODO(), 30)

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

	_, err = r.InsertNewPokemon(context.TODO(), randomPokemon)
	if err != nil {
		fmt.Println("Error when inserting pokemon")
	}

	fmt.Println("SecretPokemon Updated.")

	return err
}

func UpdateCurrentDailyStatsWithGamesWon() error {
	mongoClient := data.GetMongoDBClient()
	r := dailyStats.GetPokemonRepository(mongoClient)

	currentTime := time.Now()

	currentDate := currentTime.Format("2006-01-02")

	err := r.UpdateDailyGuessCount(context.TODO(), currentDate)

	return err
}
