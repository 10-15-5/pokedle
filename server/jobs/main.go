package jobs

import (
	"fmt"

	"github.com/gabr0236/pokedle/server/services"
	"github.com/robfig/cron"
)

func StartJobs() {
	c := cron.New()
	updateSecretPokemonJob(c)
	c.Start()
}

func updateSecretPokemonJob(c *cron.Cron) {
	c.AddFunc("@daily", func() {
		fmt.Println("Job Running.....")
		services.UpdateDailySecretPokemon()
	})
}
