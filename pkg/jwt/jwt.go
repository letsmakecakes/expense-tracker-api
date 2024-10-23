package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define your JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Secret key (loaded from environment or config)
// var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken generates a JWT for a valid user
func GenerateToken(username string, jwtKey []byte) (string, error) {
	expirationTime := time.Now().Add(12 * time.Hour) // Token valid for 12 hours

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}