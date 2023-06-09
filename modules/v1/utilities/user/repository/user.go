package repository

import (
	"fmt"
	role "rest-api/modules/v1/utilities/role/models"
	model "rest-api/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetById(id int) (model.Users, error)
	Create(user model.Users) (model.Users, error)
	Login(loginRequest model.UserLogin, refresh_token string) (*model.Users, error)
	Refresh(email string, cookie string) (*model.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) GetById(id int) (model.Users, error) {
	var user model.Users

	err := repo.db.Find(&user, id).Error

	return user, err
}

func (repo *repository) FindRoles(id int) (string, error) {
	var roleName role.Roles
	err := repo.db.Find(&roleName, id).Error

	return *roleName.Name, err
}

/* func (repo *repository) GetUser(email string) (model.Users, error) {
	var user model.Users
	err := repo.db.First(&user, "email = ?", email).Error

	return user, err
} */

func (repo *repository) Create(user model.Users) (model.Users, error) {

	err := repo.db.Create(&user).Error

	fmt.Println("repo", user)
	return user, err
}

func (repo *repository) Login(loginRequest model.UserLogin, refresh_token string) (*model.Users, error) {
	var user model.Users

	//user.Refresh_token = refresh_token

	tx := repo.db.Begin()

	defer func() {
		if repo := recover(); repo != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&user).Where("id = ?", user.ID).Update("refresh_token", refresh_token).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit().Error

}

func (repo *repository) Refresh(email string, cookie string) (*model.Users, error) {
	var user model.Users

	tx := repo.db.Begin()

	defer func() {
		if repo := recover(); repo != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.First(&user, "email = ?", email).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit().Error
}

/* func CompareString(str1, str2 string) bool {
	return str1 == str2
} */
