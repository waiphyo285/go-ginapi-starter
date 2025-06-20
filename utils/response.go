package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Respond(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Data:    data,
	})
}

func RespondOK(c *gin.Context, data interface{}) {
	Respond(c, http.StatusOK, data)
}

func RespondCreated(c *gin.Context, data interface{}) {
	Respond(c, http.StatusCreated, data)
}

func RespondError(c *gin.Context, status int, err interface{}) {
	c.JSON(status, Response{
		Success: false,
		Error:   err,
	})
}
