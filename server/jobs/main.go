package jobs

import (
	"github.com/gabr0236/pokedle/server/services"
	"github.com/robfig/cron"
)

func StartJobs() {
	c := cron.New()
	updateSecretPokemonJob(c)
	c.Start()
}

func updateSecretPokemonJob(c *cron.Cron) {
	c.AddFunc("0 00 * * *", func() {
		services.UpdateDailySecretPokemon()
	})
}
