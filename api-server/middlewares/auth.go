package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO load from db
const API_KEY = "896me27risv1r3ndp5bs074ccjacmsdc"

// AuthenticationMiddleware checks if the request has a valid api key
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("X-API-Key")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "error_message": "Missing api key"})
			c.Abort()
			return
		}

		if tokenString != API_KEY {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "error_message": "Invalid api key"})
			c.Abort()
		}

		c.Next()
	}
}
