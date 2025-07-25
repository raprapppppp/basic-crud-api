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
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(user models.Users, id int) (models.Users, error)
	DeleteUser(user models.Users, id int) error

	CreateAccountService(account *models.Account) error
	LoginUserAccountRepo(username string) (models.Account, error)
	CheckUsernameAlreadyExist(username string) bool
	GetProfileRepo(id int) (models.Account, error)

	//Task
	CreateTaskRepo(task models.Task) (models.Task, error)
	GetTaskRepo(id int) ([]models.Task, error)
	DeleteTaskRepo(task models.Task) error
	UpdateTaskRepo(task models.Task) (models.Task, error)

	 CheckEmailIfExist(user models.Users) bool
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

func (r *userDbRepo) CreateUser(user models.Users) (models.Users, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *userDbRepo) CheckEmailIfExist(user models.Users) bool {
	var count int64
	r.db.Model(&user).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
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

func (r *userDbRepo) CheckUsernameAlreadyExist(username string) bool {
	var user models.Account
	var count int64

	r.db.Model(&user).Where("Username = ?", username).Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}
//
func (r *userDbRepo) GetProfileRepo(id int) (models.Account, error) {

	var users models.Account
	err := r.db.Where("id = ?", id).Select("id","username","role").Find(&users).Error
	if err != nil {
		return models.Account{}, err
	}
	return users,nil
}

//Task
//Create
func (r *userDbRepo) CreateTaskRepo(task models.Task) (models.Task, error){
	err := r.db.Create(&task).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

//Get
func (r *userDbRepo) GetTaskRepo(id int) ([]models.Task, error){
	var task []models.Task

	err := r.db.Where("account_id = ?", id).Find(&task).Error
	if err != nil {
		return []models.Task{}, err
	}
	return task, nil
}

//delete
func (r *userDbRepo) DeleteTaskRepo(task models.Task) error {
	err := r.db.Delete(&task, task.ID).Error

	if err != nil {
		return err
	}
	return nil
}

//Update
func (r *userDbRepo) UpdateTaskRepo(task models.Task) (models.Task, error) {
	err := r.db.Where("id = ?", task.ID).Updates(&task).Error

	if err != nil {
		return models.Task{} ,err
	}
	return task, nil
}
