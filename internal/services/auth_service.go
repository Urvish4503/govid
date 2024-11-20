package services

import (
	"errors"

	"github.com/Urvish4503/govid/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Register(email, password string) (string, error) {
	var user models.User

	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email")
		}
		return "", err
	}

	hashedPassword, salt := user.Password, user.Salt

	return "", nil
}
