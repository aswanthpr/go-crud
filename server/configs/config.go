package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string, defualtValue ...string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	if len(defualtValue) > 0 {
		return defualtValue[0]
	}
	return ""
}