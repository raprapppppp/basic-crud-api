// go_fiber/services/user_service.go
package services

import (
	"go_fiber/models"

	"gorm.io/gorm"
)

// UserService defines the interface for user-related operations.
// This makes your code more testable and flexible.
type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	// Add more methods like UpdateUser, DeleteUser, ListUsers etc.
}

// userService implements UserService using GORM.
type userService struct {
	db *gorm.DB
}

// NewUserService creates a new instance of UserService.
// It takes the GORM DB client as a dependency.
func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

// CreateUser handles the creation of a new user in the database.
func (s *userService) CreateUser(user *models.User) error {
	// GORM's Create method directly saves the user and updates the user.ID
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	// No need to set user.ID here, GORM does it automatically if primaryKey is setup.
	return nil
}

// GetUserByID retrieves a user by their ID.
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	result := s.db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
