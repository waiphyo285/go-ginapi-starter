package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"neohub.asia/mod/controllers"
	"neohub.asia/mod/middlewares"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Inject DB into context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Add Rate Limit
	r.Use(middlewares.NewRateLimiter(5, 10*time.Second).Middleware())

	// Public routes
	r.POST("/auth/token", controllers.LoginHandler)

	// Protected routes
	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	api.Use(middlewares.ResponseFormatter())

	// Register route groups
	RegisterBookRoutes(api, db)

	return r
}
