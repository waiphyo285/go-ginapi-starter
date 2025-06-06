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
		api.GET("/book", controllers.FindTasks)
		api.POST("/book", controllers.CreateTask)
		api.GET("/book/:id", controllers.FindTask)
		api.PATCH("/book/:id", controllers.UpdateTask)
		api.DELETE("/book/:id", controllers.DeleteTask)
	}
	return r
}
