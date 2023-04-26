package repository

import (
	model "rest-api/modules/v1/utilities/user/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByID(ID int) (model.Membership, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (model.Membership, error) {
	var membership model.Membership
	err := r.db.Find(&membership, ID).Error
	return membership, err
}
