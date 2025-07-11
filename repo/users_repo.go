package repo

import (
	"errors"
	"go_fiber/models"

	"gorm.io/gorm"
)

// Interface contains available method in UserRepo
type UserRepo interface {
	FindAll() ([]models.Users, error)
	//FindByID(id int) (models.User, error)
	CreateUser(user *models.Users) (models.Users, error)
	UpdateUser(user models.Users, id int) (models.Users, error)
	DeleteUser(user models.Users, id int) error

	CreateAccountService(account *models.Account) error
	LoginUserAccountRepo(username string) (models.Account, error)
}

// Inject DB
type userDbRepo struct {
	db *gorm.DB
}

// to initialize
func NewUserRepository(db *gorm.DB) UserRepo {
	return &userDbRepo{db}
}

// getter
func (r *userDbRepo) FindAll() ([]models.Users, error) {
	var user []models.Users

	err := r.db.Order("id asc").Find(&user).Error
	return user, err
}

func (r *userDbRepo) CreateUser(user *models.Users) (models.Users, error) {

	err := r.db.Create(&user).Error

	return *user, err
}

func (r *userDbRepo) UpdateUser(user models.Users, id int) (models.Users, error) {

	err := r.db.Where("id = ?", id).Updates(&user).Error

	return user, err
}

func (r *userDbRepo) DeleteUser(user models.Users, id int) error {

	err := r.db.Delete(&user, id).Error

	return err
}

func (r *userDbRepo) CreateAccountService(account *models.Account) error {

	err := r.db.Create(&account).Error

	if err != nil {
		return errors.New("username already exist")
	}
	return err
}

func (r *userDbRepo) LoginUserAccountRepo(account string) (models.Account, error) {
	var accounts models.Account

	var notFoundError = errors.New("user not found")

	err := r.db.Find(&accounts, "username = ?", account).Error
	if err != nil {
		return models.Account{}, notFoundError
	}
	return accounts, nil
}
