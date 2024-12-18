package routes

import (
	"log"

	"github.com/Urvish4503/govid/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPingRouter(app *fiber.App) {
	pingRouter := app.Group("/api")

	pingRouter.Get("/ping", func(c *fiber.Ctx) error {
		log.Println("Ping")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "pong",
		})
	})
}

func SetupAuthRouter(app *fiber.App, handlers *handlers.AuthHandler) {
	authRouter := app.Group("/auth")

	authRouter.Post("/register", handlers.Register)
	authRouter.Post("/login", handlers.Login)

}

func SetupUserRouter(app *fiber.App, handlers *handlers.UserHandler) {
	userRouter := app.Group("/user")

	userRouter.Get("/", handlers.GetUser)
	userRouter.Put("/", handlers.UpdateUser)
	userRouter.Delete("/", handlers.DeleteUser)
}
