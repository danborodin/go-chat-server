package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVar func return env value
func GetEnvVar(key string) string {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
