package handlers

import (
	"github.com/Urvish4503/govid/internal/models"
	"github.com/Urvish4503/govid/internal/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var loginData models.LoginRequest

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong input",
		})
	}

	return nil
}
