package handlers

import (
	"time"

	"github.com/Urvish4503/govid/internal/models"
	"github.com/Urvish4503/govid/internal/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var userReq models.RegisterRequest

	// Parse the request body
	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Register the user
	user, err := h.authService.Register(&userReq)

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

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginData models.LoginRequest

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong input",
		})
	}

	// Generating JWT token
	token, err := h.authService.Login(loginData.Email, loginData.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
			"error":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}
