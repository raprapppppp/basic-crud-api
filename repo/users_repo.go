package repo

import (
	"go_fiber/models"

	"gorm.io/gorm"
)

// Interface contains available method in UserRepo
type UserRepo interface {
	FindAll() ([]models.User, error)
	//FindByID(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, id int) (models.User, error)
	DeleteUser(user models.User, id int) error
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
func (r *userDbRepo) FindAll() ([]models.User, error) {
	var user []models.User

	err := r.db.Find(&user).Error
	return user, err
}

func (r *userDbRepo) CreateUser(user models.User) (models.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *userDbRepo) UpdateUser(user models.User, id int) (models.User, error) {

	err := r.db.Where("id = ?", id).Updates(&user).Error

	return user, err
}

func (r *userDbRepo) DeleteUser(user models.User, id int) error {

	err := r.db.Delete(&user, id).Error

	return err
}
