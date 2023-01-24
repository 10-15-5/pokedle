package jobs

import (
	"github.com/robfig/cron"
)

func StartJobs() {
	c := cron.New()
	updateSecretPokemonJob(c)
	c.Start()
}

func updateSecretPokemonJob(c *cron.Cron) {
	c.AddFunc("@every 5s", func() {
		updateDailySecretPokemon()
	})
}
