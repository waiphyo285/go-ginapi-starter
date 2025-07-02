package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"neohub.asia/mod/controllers"
	"neohub.asia/mod/di"
	"neohub.asia/mod/middlewares"
)

func SetupRoutes(c *di.Container) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Inject DB into context
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", c.DB)
	})

	// Add Rate Limit
	r.Use(middlewares.NewRateLimiter(10, 10*time.Second).Middleware())

	// Public routes
	r.POST("/auth/token", controllers.LoginHandler)

	// Protected routes
	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	api.Use(middlewares.ResponseFormatter())

	// Register route groups
	RegisterBookRoutes(api, c)

	return r
}
