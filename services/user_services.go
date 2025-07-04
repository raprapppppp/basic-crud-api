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
	FindAll() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, id int) (models.User, error)
	DeleteUser(user models.User, id int) error
}

// Init
func UserServiceInit(r repo.UserRepo) UserServiceDepend {
	return &UserService{r}
}

func (s *UserService) FindAll() ([]models.User, error) {
	return s.service.FindAll()
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.service.CreateUser(user)
}

func (s *UserService) UpdateUser(user models.User, id int) (models.User, error) {
	return s.service.UpdateUser(user, id)
}

func (s *UserService) DeleteUser(user models.User, id int) error {
	return s.service.DeleteUser(user, id)
}
