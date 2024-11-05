package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	Environment string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Environment: os.Getenv("ENVIRONMENT"),
	}, nil
}
