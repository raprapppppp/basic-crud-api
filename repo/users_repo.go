package repo

import (
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

func (r *userDbRepo) UpdateUser(user models.Users, id int) (models.Users, error) {

	err := r.db.Where("id = ?", id).Updates(&user).Error

	return user, err
}

func (r *userDbRepo) DeleteUser(user models.Users, id int) error {

	err := r.db.Delete(&user, id).Error

	return err
}
