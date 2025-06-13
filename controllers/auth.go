package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"neohub.asia/mod/config"
	"neohub.asia/mod/services/jwt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func LoginHandler(c *gin.Context) {
	var body LoginRequest
	jwtCfg := config.LoadJWTConfig()
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if body.Username != "admin" || body.Password != "secret" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := jwtservice.CreateToken(
		map[string]interface{}{"sub": body.Username},
		time.Minute*time.Duration(jwtCfg.ExpiresIn),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken: token,
		TokenType:   "bearer",
	})
}
