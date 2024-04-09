package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	Logger.Info("Loading environment variables...")
	err := godotenv.Load(".env")
	if err != nil {
		Logger.Fatalf("Failed to load environment variables: %v", err)
	}
	Logger.Info("Environment variables loaded successfully")
}

func GetEnv(key string) string {
	Logger.Debugf("Getting value for environment variable '%s'", key)
	value, exists := os.LookupEnv(key)
	if !exists {
		Logger.Warnf("Environment variable %s is not set", key)
	}
	Logger.Debugf("Value for environment variable '%s': %s", key, value)
	return value
}

func MustGetEnv(key string) string {
	Logger.Debugf("Getting value for environment variable '%s'", key)
	value, exists := os.LookupEnv(key)
	if !exists {
		Logger.Fatalf("Environment variable %s is not set", key)
	}
	Logger.Debugf("Value for environment variable '%s': %s", key, value)
	return value
}
