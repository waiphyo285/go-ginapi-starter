package routes

import (
	"github.com/gin-gonic/gin"
	"neohub.asia/mod/di"
)

func RegisterBookRoutes(rg *gin.RouterGroup, c *di.Container) {
	bookController := c.BookController

	book := rg.Group("/book")
	{
		book.GET("/", bookController.List)
		book.POST("/", bookController.Create)
		book.GET("/:id", bookController.Get)
		book.PATCH("/:id", bookController.Update)
		book.DELETE("/:id", bookController.Delete)
	}
}
