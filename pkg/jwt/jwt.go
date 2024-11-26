package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims Define your JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Secret key (loaded from environment or config)
var jwtKey = []byte("your_secret_key")

// GenerateToken generates a JWT for a valid user
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(12 * time.Hour) // Set token expiration time to 12 hours

	// Create the JWT claims, which includes the username and expiry time
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

// ValidateToken validates the incoming JWT and returns the claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, fmt.Errorf("invalid signature")
		}
		return nil, fmt.Errorf("could not parse token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}
