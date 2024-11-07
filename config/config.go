package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseURL string
	Environment string
}

func LoadConfig() (*Config, error) {
	// Set the path to look for the .env file
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading config file: %w", err)
	}

	// Set default values (if needed)
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENVIRONMENT", "development")

	// Bind environment variables to Viper
	viper.AutomaticEnv()

	return &Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		Environment: viper.GetString("ENVIRONMENT"),
	}, nil
}
