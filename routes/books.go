package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"neohub.asia/mod/controllers"
	"neohub.asia/mod/databases/models"
)

func RegisterBookRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	bookController := controllers.NewBaseController(db, models.Book{})
	book := rg.Group("/book")
	{
		book.GET("/", bookController.List)
		book.POST("/", bookController.Create)
		book.GET("/:id", bookController.Get)
		book.PATCH("/:id", bookController.Update)
		book.DELETE("/:id", bookController.Delete)
	}
}
