package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"neohub.asia/mod/controllers"
	"neohub.asia/mod/middlewares"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Public route 
	r.POST("/auth/token", controllers.LoginHandler)

	// Protected routes
	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	{
		api.GET("/book", controllers.GetBooks)
		api.POST("/book", controllers.CreateBook)
		api.GET("/book/:id", controllers.GetBook)
		api.PATCH("/book/:id", controllers.UpdateBook)
		api.DELETE("/book/:id", controllers.DeleteBook)
	}
	return r
}
