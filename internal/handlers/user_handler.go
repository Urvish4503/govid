package handlers

import (
	"github.com/Urvish4503/govid/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	return nil
}

// TODO: Implement GetUser handler
func (h *UserHandler) GetUser(c *fiber.Ctx) error {

	return nil
}
