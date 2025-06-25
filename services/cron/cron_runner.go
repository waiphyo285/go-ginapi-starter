package cronservice

import (
	"log"

	"github.com/robfig/cron/v3"
)

var (
	Cron    *cron.Cron
	jobChan chan string
)

func StartWorker() {
	go func() {
		for msg := range jobChan {
			log.Println("Processing job message:", msg)
		}
	}()
}

func CronRunner() {
	Cron = cron.New()
	jobChan = make(chan string, 100)

	StartWorker()

	_, err := Cron.AddFunc("0 2 * * *", SayGreetingJob)
	if err != nil {
		panic("Failed to schedule SayGreetingJob: " + err.Error())
	}

	Cron.Start()
}
