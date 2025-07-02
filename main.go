package main

import (
	"neohub.asia/mod/di"
	"neohub.asia/mod/routes"
	cronservice "neohub.asia/mod/services/cron"
)

func main() {
	container := di.NewContainer()

	// Inject DB into cron jobs
	cronservice.CronRunner()

	r := routes.SetupRoutes(container)
	r.Run(":9002")
}
