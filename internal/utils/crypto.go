package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword combines password with salt and hashes
func HashPassword(password string) (string, error) {
	saltedPassword := password

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedBytes), nil
}

// Validate password
func ValidatePassword(providedPlainPassword, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPlainPassword))

	return err == nil

}
