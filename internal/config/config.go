package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/m-golang/weather-app/internal/helpers"
)

// LoadEnvFile loads the environment variables from the .env file
func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// LoadAPIKey retrieves the API key for the weather service from the environment variables.
func LoadAPIKey(key string) (string, error) {
	apiKey := os.Getenv(key)
	if apiKey == "" {
		return "", helpers.ErrAPIKeyNotFound
	}
	return apiKey, nil
}
