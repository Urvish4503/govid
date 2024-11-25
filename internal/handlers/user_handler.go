package handlers

import (
	"github.com/Urvish4503/govid/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	GetUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	token := c.Cookies("jwt")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	user, err := h.userService.GetUser(token)

	return nil
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	return nil
}
