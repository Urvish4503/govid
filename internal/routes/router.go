package routes

import (
	"github.com/Urvish4503/govid/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPingRouter(app *fiber.App) {
	pingRouter := app.Group("/api")

	pingRouter.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "pong",
		})
	})
}

func SetupUserRouter(app *fiber.App, handlers *handlers.UserHandler) {
	userRouter := app.Group("/user")

	userRouter.Post("/register", handlers.RegisterUser)
}
