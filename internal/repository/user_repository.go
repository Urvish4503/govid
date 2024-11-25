package repository

import (
	"errors"

	"github.com/Urvish4503/govid/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(email string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(userID uuid.UUID) error
}

type PostgresUserRepository struct {
	db *gorm.DB
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserUpdate   = errors.New("failed to update user")
	ErrUserDelete   = errors.New("failed to delete user")
	ErrEmailExists  = errors.New("email already exists")
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) CreateUser(user *models.User) (*models.User, error) {
	existingUser, err := r.GetUser(user.Email)

	if err == nil && existingUser != nil {
		return nil, ErrEmailExists
	}
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return nil, err // Return any other unexpected error
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) GetUser(email string) (*models.User, error) {
	user := new(models.User)

	result := r.db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, result.Error
	}

	return user, nil
}

func (r *PostgresUserRepository) UpdateUser(user *models.User) (*models.User, error) {
	result := r.db.Save(user)

	if result.Error != nil {
		return nil, ErrUserUpdate
	}

	return user, nil
}

func (r *PostgresUserRepository) DeleteUser(userID uuid.UUID) error {
	result := r.db.Delete(&models.User{}, userID)

	if result.Error != nil {
		return ErrUserDelete
	}

	return nil
}
