package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf(".env.%s", env)

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("⚠️  Warning: could not load env file: %s (%v)\n", envFile, err)
	} else {
		log.Printf("✅ Loaded environment: %s\n", envFile)
	}
}

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
