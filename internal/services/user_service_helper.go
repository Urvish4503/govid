package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

	"github.com/Urvish4503/govid/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// GenerateSalt creates a random salt
func generateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// HashPassword combines password with salt and hashes
func hashPassword(password string, salt string) (string, error) {
	saltedPassword := password + salt

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedBytes), nil
}

func validateUserRequest(user *models.UserRequest) error {
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
