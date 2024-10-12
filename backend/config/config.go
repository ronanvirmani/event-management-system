package config

import (
	"log"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}