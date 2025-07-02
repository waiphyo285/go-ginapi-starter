package di

import (
	"gorm.io/gorm"
	"neohub.asia/mod/controllers"
	"neohub.asia/mod/databases"
	"neohub.asia/mod/databases/models"
)

type Container struct {
	DB             *gorm.DB
	BookController *controllers.BaseController[models.Book]
}

func NewContainer() *Container {
	// Initialize Db connection
	db := databases.SetupDB()

	// Migrate models
	db.AutoMigrate(&models.Book{}, &models.AuditLog{})

	// Register Hooks
	databases.RegisterHooks(db)

	return &Container{
		DB:             db,
		BookController: controllers.NewBaseController(db, models.Book{}),
	}
}
