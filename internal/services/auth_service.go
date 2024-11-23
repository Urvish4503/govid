package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Urvish4503/govid/internal/models"
	"github.com/Urvish4503/govid/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Register(userReq *models.UserRequest) (*models.User, error) {

	if err := utils.ValidateUserRequest(userReq); err != nil {
		return nil, err
	}

	var existingUser models.User

	result := s.DB.Where("email = ?", userReq.Email).First(&existingUser)

	if result.Error == nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(userReq.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Name:     userReq.Name,
		Email:    strings.ToLower(userReq.Email),
		Password: hashedPassword,
	}

	if err := s.DB.Create(user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email")
		}
		return "", err
	}

	fmt.Println(user)

	hashedPassword := user.Password

	if !utils.ValidatePassword(password, hashedPassword) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.Email, user.Name)

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
