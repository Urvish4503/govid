package main

import (
	"log"

	"github.com/Urvish4503/govid/internal/config"
	"github.com/Urvish4503/govid/internal/models"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

}

func main() {
	err := config.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed")
}
