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
		// Skip authentication for POST /signup
		if c.Request.Method == http.MethodPost && c.FullPath() == "/signup" {
			c.Next()
			return
		}

		// Skip authentication for POST /login
		if c.Request.Method == http.MethodPost && c.FullPath() == "/login" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		// Retrieve the JWT token from the cookie
		//tokenString, err := c.Cookie("token")
		//if err != nil {
		//	utils.RespondWithError(c, http.StatusUnauthorized, "Authorization token not found")
		//	c.Abort()
		//	return
		//}

		// Retrieve the token from header
		tokenString := authHeader[7:]

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
