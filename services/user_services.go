// go_fiber/services/user_service.go
package services

import (
	"go_fiber/models"
	"go_fiber/repo"
)

// injcting UserRepo interfaces
type UserService struct {
	service repo.UserRepo
}

// interfaces contains all mthod available on services
type UserServiceDepend interface {
	FindAll() ([]models.Users, error)
	CreateUser(user *models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(user models.Users) error
}

// Init
func UserServiceInit(r repo.UserRepo) UserServiceDepend {
	return &UserService{r}
}

func (s *UserService) FindAll() ([]models.Users, error) {
	return s.service.FindAll()
}

func (s *UserService) CreateUser(user *models.Users) (models.Users, error) {

	return s.service.CreateUser(user)
}

func (s *UserService) UpdateUser(user models.Users) (models.Users, error) {

	id := user.ID
	return s.service.UpdateUser(user, int(id))
}

func (s *UserService) DeleteUser(user models.Users) error {

	id := user.ID
	return s.service.DeleteUser(user, int(id))
}
