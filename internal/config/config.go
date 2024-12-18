package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Config() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitDB()
}
