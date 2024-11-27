package middleware

import (
	"expensetrackerapi/pkg/jwt"
	"expensetrackerapi/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware Middleware to authenticate JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header is required and must contain a Bearer token")
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenString := authHeader[7:]

		// Parse and validate the token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Store user info in context for use in subsequent handlers
		c.Set("username", claims.Username)

		// Continue to the next handler
		c.Next()
	}
}
