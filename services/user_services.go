// go_fiber/services/user_service.go
package services

import (
	"fmt"
	"go_fiber/crypt/util"
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

	CreateAccountService(account *models.Account) (string, error)
	LoginUserAccountService(loginCredential models.Account) (models.Account, string)
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

func (s *UserService) CreateAccountService(account *models.Account) (string, error) {

	fmt.Print(account)

	isExist := s.service.CheckUsernameAlreadyExist(account.Username)
	if isExist {
		return "Exist",nil
	}

	account.Password = util.HashPassword(account.Password)
	err := s.service.CreateAccountService(account)
	if err != nil {
		return "",err
	}
	return "Created", nil
}

func (s *UserService) LoginUserAccountService(loginCredential models.Account) (models.Account, string) {
	
	isAlreadyExist := s.service.CheckUsernameAlreadyExist(loginCredential.Username)
	if !isAlreadyExist {
		return models.Account{}, "User Does not exist exist"
	}

	acc, err := s.service.LoginUserAccountRepo(loginCredential.Username)
	if err != nil {
		return models.Account{}, err.Error()
	}

	isMatch := util.CompareHashAndPassword(acc.Password, loginCredential.Password)
	if !isMatch {
		return models.Account{}, "Password does not match"
	}
	return acc, "Account match"
}
