package main

import (
	"neohub.asia/mod/databases/models"
	"neohub.asia/mod/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Book{})

	r := routes.SetupRoutes(db)
	r.Run(":9002")
}
