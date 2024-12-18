package services

import (
	"github.com/Urvish4503/govid/internal/models"
	"github.com/Urvish4503/govid/internal/repository"
)

type UserServiceInterface interface {
	GetUser(email string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(email string) error
}

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUser(email string) (*models.User, error) {
	// TODO:
	user, err := s.userRepo.GetUser(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {

	user, err := s.userRepo.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) DeleteUser(email string) error {
	user, err := s.userRepo.GetUser(email)

	if err != nil {
		return err
	}

	err = s.userRepo.DeleteUser(user.ID)

	return err

}
