package utils

import (
	"regexp"
	"strings"

	"errors"
	"github.com/Urvish4503/govid/internal/models"
)

func ValidateUserRequest(user *models.RegisterRequest) error {
	// Validate name
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name is required")
	}
	if len(user.Name) > 255 {
		return errors.New("name is too long")
	}

	// Validate email
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(strings.ToLower(user.Email)) {
		return errors.New("invalid email format")
	}

	// Validate password
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if len(user.Password) > 255 {
		return errors.New("password is too long")
	}

	return nil
}
