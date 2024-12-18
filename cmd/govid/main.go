package main

import (
	"log"

	"github.com/Urvish4503/govid/internal/config"
	"github.com/Urvish4503/govid/internal/handlers"
	"github.com/Urvish4503/govid/internal/repository"
	"github.com/Urvish4503/govid/internal/routes"
	"github.com/Urvish4503/govid/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.Config()
}

func main() {
	// userService := services.NewUserService(config.DB)
	// userHandler := handlers.NewUserHandler(userService)
	//
	// authService := services.NewAuthService(config.DB)
	// authHandler := handlers.NewAuthHandler(authService)

	userRepo := repository.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)


	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
	}))

	routes.SetupPingRouter(app)

	routes.SetupAuthRouter(app, authHandler)

	// routes.SetupAuthRouter(app, authHandler)
	// routes.SetupUserRouter(app, userHandler)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
