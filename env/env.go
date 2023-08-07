package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetAPIKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	return os.Getenv("API_KEY")
}
