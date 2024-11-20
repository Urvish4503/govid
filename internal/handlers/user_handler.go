package handlers

import (
	"github.com/Urvish4503/govid/internal/models"
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
	var userReq models.UserRequest

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user, err := h.userService.RegisterUser(&userReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    user,
	})
}

// TODO: Implement GetUser handler
func (h *UserHandler) GetUser(c *fiber.Ctx) error {

	return nil
}
