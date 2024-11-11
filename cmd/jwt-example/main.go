package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// Secret key for signing the token (in a real application, store this securely)
var jwtKey = []byte("your_secret_key")

// Claims represents the structure of the JWT payload
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT for a specific user
func GenerateJWT(username string) (string, error) {
	// Set token expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a JWT and extracts the username from it
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func main() {
	// Generate a JWT for a specific user
	token, err := GenerateJWT("adwaith")
	if err != nil {
		log.Fatalf("Error generating token: %v", err)
	}

	fmt.Printf("Generated Token: %s\n", token)

	// Validate the JWT and extract claims
	claims, err := ValidateJWT(token)
	if err != nil {
		log.Fatalf("Error validating token: %v", err)
	}

	fmt.Printf("Token is valid. Username: %s\n", claims.Username)
}
