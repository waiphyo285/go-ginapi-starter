package cronservice

import (
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func CronRunner() {
	Cron = cron.New()

	// Schedule: Every day at 02:00 AM
	_, err := Cron.AddFunc("0 2 * * *", SayGreetingJob)
	if err != nil {
		panic("Failed to schedule SayGreetingJob: " + err.Error())
	}

	Cron.Start()
}
