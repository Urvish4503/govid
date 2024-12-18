package services

import (
	"errors"
	"strings"

	"github.com/Urvish4503/govid/internal/models"
	"github.com/Urvish4503/govid/internal/repository"
	"github.com/Urvish4503/govid/internal/utils"
	"github.com/google/uuid"
)

// AuthService interface defines the contract for authentication operations
type AuthServiceInterface interface {
	Register(request *models.RegisterRequest) (*models.User, error)
	Login(email, password string) (string, error)
}

// UserAuthService implements AuthService interface
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Register handles user registration
func (s *AuthService) Register(userReq *models.RegisterRequest) (*models.User, error) {
	if err := utils.ValidateUserRequest(userReq); err != nil {
		return nil, err
	}

	// Check if email already exists
	existingUser, err := s.userRepo.GetUser(userReq.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := &models.User{
		ID:       uuid.New(),
		Name:     userReq.Name,
		Email:    strings.ToLower(userReq.Email),
		Password: hashedPassword,
	}

	// Save user using repository
	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return createdUser, nil
}

// Login handles user authentication
func (s *AuthService) Login(email, password string) (string, error) {
	// Get user by email
	user, err := s.userRepo.GetUser(email)

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	// Validate password
	if !utils.ValidatePassword(password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Email, user.Name)
	if err != nil {
		return "", ErrTokenGeneration
	}

	return token, nil
}

// Common errors
var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailExists        = errors.New("email already registered")
	ErrTokenGeneration    = errors.New("failed to generate token")
)
