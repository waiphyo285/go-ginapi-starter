package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtservice "neohub.asia/mod/services/jwt"
	"neohub.asia/mod/utils"
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
	
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if body.Username != "admin" || body.Password != "secret" {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid username or password!")
		return
	}


	token, err := jwtservice.CreateToken(
		map[string]interface{}{"sub": body.Username},
	)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create token!")
		return
	}

	utils.RespondOK(c, TokenResponse{
		AccessToken: token,
		TokenType:   "bearer",
	})
}
