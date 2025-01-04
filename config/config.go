package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config holds the configuration values for the application.
type Config struct {
	Port        string
	DatabaseURL string
	Environment string
}

// LoadConfig initializes and loads the configuration from file and environment variables.
func LoadConfig() (*Config, error) {
	if err := setupViper(); err != nil {
		return nil, fmt.Errorf("error setting up Viper: %w", err)
	}

	if err := readConfigFile(); err != nil {
		return nil, fmt.Errorf("error loading config file: %w", err)
	}

	setDefaults()
	viper.AutomaticEnv()

	return &Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		Environment: viper.GetString("ENVIRONMENT"),
	}, nil
}

// setupViper configures Viper with the paths and filename of the configuration file.
func setupViper() error {
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../")
	viper.SetConfigName(".env")
	return nil
}

// readConfigFile reads the configuration file using Viper.
func readConfigFile() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// setDefaults sets the default values for the configuration.
func setDefaults() {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENVIRONMENT", "development")
}
