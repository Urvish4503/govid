package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSalt creates a random salt
func GenerateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// HashPassword combines password with salt and hashes
func HashPassword(password, salt string) (string, error) {
	saltedPassword := password + salt

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedBytes), nil
}

// Validate password
func ValidatePassword(providedPlainPassword, hashedPassword, salt string) bool {

	hashOfProvidedPassword, err := HashPassword(providedPlainPassword, salt)

	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(hashOfProvidedPassword), []byte(hashedPassword)) == 1
}
