package main

import (
	"log"

	"github.com/Urvish4503/govid/internal/config"
	"github.com/Urvish4503/govid/internal/handlers"
	"github.com/Urvish4503/govid/internal/routes"
	"github.com/Urvish4503/govid/internal/services"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.Config()
}

func main() {
	// userService := services.NewUserService(config.DB)
	// userHandler := handlers.NewUserHandler(userService)

	authService := services.NewAuthService(config.DB)
	authHandler := handlers.NewAuthHandler(authService)

	app := fiber.New()

	routes.SetupPingRouter(app)

	routes.SetupAuthRouter(app, authHandler)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
