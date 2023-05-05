package repository

import (
	//model "rest-api/modules/v1/utilities/user/model"

	"gorm.io/gorm"
)

/* type IUserRepository interface {

} */

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}
