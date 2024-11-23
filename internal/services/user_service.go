package services

import (
	"github.com/Urvish4503/govid/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}


func (s *UserService) GetUser(email string) (*models.User, error) {
	return nil, nil
}
