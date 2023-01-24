package jobs

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gabr0236/pokedle/dev-server/database"
	"github.com/gabr0236/pokedle/dev-server/models"
	"github.com/gabr0236/pokedle/dev-server/repository"
)

func updateDailySecretPokemon() {
	fmt.Println("Enter updateDailySecretPokemon")

	pokemonData := database.GetPokemonData()

	mongoClient := database.GetMongoDBClient()

	if pokemonData == nil || mongoClient == nil {
		panic("Missing pokemonData or mongoClient")
	}

	r := repository.GetRepository(mongoClient)
	fmt.Println("GetRepository Successfull")

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

	result, err := r.InsertNewPokemon(context.TODO(), randomPokemon)
	if err != nil {
		fmt.Println("Error when inserting pokemon")
	}

	fmt.Println("Pokemon inserted. DocumentId is:", result)
}
