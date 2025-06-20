package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"neohub.asia/mod/utils"
)

func ResponseFormatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Written() {
			return
		}

		// Handle errors
		if errResp, exists := c.Get("error"); exists {
			statusCode := http.StatusBadRequest
			var errData interface{} = errResp

			if errMap, ok := errResp.(map[string]interface{}); ok {
				if code, ok := errMap["code"].(int); ok {
					statusCode = code
				}
				if data, ok := errMap["data"]; ok {
					errData = data
				}
			}

			utils.RespondError(c, statusCode, errData)
			return
		}

		// Handle success
		if dataResp, exists := c.Get("response"); exists {
			statusCode := http.StatusOK

			if code, exists := c.Get("status"); exists {
				if cInt, ok := code.(int); ok {
					statusCode = cInt
				}
			}

			utils.Respond(c, statusCode, dataResp)
			return
		}
	}
}
