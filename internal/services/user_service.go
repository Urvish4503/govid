package services

import (
	"errors"
	"strings"

	"github.com/Urvish4503/govid/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) RegisterUser(userReq *models.UserRequest) (*models.User, error) {
	if err := validateUserRequest(userReq); err != nil {
		return nil, err
	}

	var existingUser models.User

	result := s.db.Where("email = ?", userReq.Email).First(&existingUser)

	// FIXME: here if email is not existing then it is printing error
	if result.Error == nil {
		return nil, errors.New("email already registered")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	salt, err := generateSalt(16)
	if err != nil {
		return nil, errors.New("failed to generate salt")
	}

	hashedPassword, err := hashPassword(userReq.Password, salt)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Name:     userReq.Name,
		Email:    strings.ToLower(userReq.Email),
		Password: hashedPassword,
		Salt:     salt,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}
