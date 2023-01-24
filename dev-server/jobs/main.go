package jobs

import (
	"github.com/gabr0236/pokedle/dev-server/services"
	"github.com/robfig/cron"
)

func StartJobs() {
	c := cron.New()
	updateSecretPokemonJob(c)
	c.Start()
}

func updateSecretPokemonJob(c *cron.Cron) {
	c.AddFunc("@every 1h", func() {
		services.UpdateDailySecretPokemon()
	})
}
