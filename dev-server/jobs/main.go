package jobs

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gabr0236/pokedle/dev-server/database"
	"github.com/gabr0236/pokedle/dev-server/repository"
	"github.com/robfig/cron"
)

func StartJobs() {
	c := cron.New()
	updateSecretPokemonJob(c)
	c.Start()
}

func updateSecretPokemonJob(c *cron.Cron) {
	c.AddFunc("@every 5s", func() {
		fmt.Println("Enter job")
		pokemonData := database.GetPokemonData()
		randomPokemon := pokemonData[rand.Intn(len(pokemonData))]
		fmt.Println(randomPokemon)

		mongoClient := database.GetMongoDBClient()

		r := repository.GetRepository(mongoClient)
		fmt.Println("GetRepository Successfull")

		result, err := r.InsertNewPokemon(context.TODO(), randomPokemon)
		if err != nil {
			fmt.Println("Error when inserting pokemon")
		}

		fmt.Println("Pokemon inserted. DocumentId is:", result)
	})
}
