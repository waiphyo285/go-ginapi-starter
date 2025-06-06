package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"neohub.asia/mod/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	api := r.Group("/api")
	{
		api.GET("/book", controllers.GetBooks)
		api.POST("/book", controllers.CreateBook)
		api.GET("/book/:id", controllers.GetBook)
		api.PATCH("/book/:id", controllers.UpdateBook)
		api.DELETE("/book/:id", controllers.DeleteBook)
	}
	return r
}
