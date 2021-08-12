package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func GetEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func TimeStamp(date string) int64 {
	time.
}
