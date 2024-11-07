package middleware

import (
	"expensetrackerapi/pkg/jwt"
	"expensetrackerapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware Middleware to authenticate JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip authentication for POST /signup
		if c.Request.Method == http.MethodPost && c.FullPath() == "/signup" {
			c.Next()
			return
		}

		// Retrieve the JWT token from the cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization token not found")
			c.Abort()
			return
		}

		// Parse and validate the token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Store user info in context for later use
		c.Set("username", claims.Username)
		c.Next()
	}
}
