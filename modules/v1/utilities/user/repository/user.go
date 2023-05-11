package repository

import (
	"fmt"
	model "rest-api/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetById(id int) (model.Users, error)
	Create(user model.Users) (model.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) GetById(id int) (model.Users, error) {
	var user model.Users
	fmt.Println("repo", id)
	err := repo.db.Find(&user, id).Error

	return user, err
}

/* func (repo *repository) GetUser(email string) (model.Users, error) {

} */

func (repo *repository) Create(user model.Users) (model.Users, error) {
	err := repo.db.Create(&user).Error

	fmt.Println("repo", user)
	return user, err
}

/* func (repo *repository) Login(user model.UserLogin) (model.Users, error) {

} */
