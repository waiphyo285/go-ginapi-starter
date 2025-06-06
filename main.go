package main

import (
	"github.com/waiphyo285/go-ginapi-starter/app/databases/models"
	"github.com/waiphyo285/go-ginapi-starter/app/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Book{})

	r := routes.SetupRoutes(db)
	r.Run(":9002")
}
