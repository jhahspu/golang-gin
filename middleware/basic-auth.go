package middleware

import "github.com/gin-gonic/gin"

// BasicAuth provided by gin for basic authorization
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		// authorize a user with username : password
		"test": "test",
	})
}
