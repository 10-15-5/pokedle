package jobs

import (
	"fmt"
	"math/rand"

	"github.com/gabr0236/pokedle/dev-server/database"
	"github.com/robfig/cron"
)

func StartJobs() {
	c := cron.New()
	updateSecretPokemonJob(c)

	//stops scheduler, does not stop running jobs
	c.Start()
}

func updateSecretPokemonJob(c *cron.Cron) {
	c.AddFunc("@every 3s", func() {
		fmt.Println("Enter job")
		pokemonData := database.GetPokemonData()
		fmt.Println(pokemonData[rand.Intn(len(pokemonData))])
	}) //47:52
}
