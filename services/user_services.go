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

	CreateAccountService(account *models.Account) error
	LoginUserAccountService(loginCredential models.Account) (bool, error)
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

func (s *UserService) CreateAccountService(account *models.Account) error {

	fmt.Print(account)

	passToHash := account.Password
	fmt.Print(passToHash)
	//Pass user input password to encrypt
	account.Password = util.HashPassword(passToHash)
	fmt.Print(account)

	return s.service.CreateAccountService(account)
}

func (s *UserService) LoginUserAccountService(loginCredential models.Account) (bool, error) {
	//Details from table
	userAccounts, _ := s.service.LoginUserAccountRepo()

	//Username and Password From input
	uname := loginCredential.Username
	pword := loginCredential.Password

	var match bool
	//Loop to check if the user INPUT match the login credentials on the table
	for _, accts := range userAccounts {
		if util.CompareHashAndPassword(accts.Password, pword) && accts.Username == uname {
			match = true

		} else {
			match = false
		}
	}
	return match, nil
}
