package main

import (
	"neohub.asia/mod/databases"
	"neohub.asia/mod/databases/models"
	"neohub.asia/mod/routes"

	cronservice "neohub.asia/mod/services/cron"
)

func main() {
	db := databases.SetupDB()
	// Migrate models
	db.AutoMigrate(&models.Book{}, &models.AuditLog{})

	// Register Hooks
	databases.RegisterHooks(db)

	// Start Runner Jobs
	cronservice.CronRunner()

	// Setup Routes
	r := routes.SetupRoutes(db)
	r.Run(":9002")
}
