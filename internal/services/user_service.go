package services

import (
	"errors"
	"github.com/Urvish4503/govid/internal/utils"
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
	if err := utils.ValidateUserRequest(userReq); err != nil {
		return nil, err
	}

	var existingUser models.User

	result := s.db.Where("email = ?", userReq.Email).First(&existingUser)

	if result.Error == nil {
		return nil, errors.New("email already registered")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
	} else {
		return nil, result.Error
	}

	salt, err := utils.GenerateSalt(16)
	if err != nil {
		return nil, errors.New("failed to generate salt")
	}

	hashedPassword, err := utils.HashPassword(userReq.Password, salt)

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

// TODO: later implement GetUser
func (s *UserService) GetUser() {}
