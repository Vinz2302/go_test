package repository

import (
	model "rest-api/modules/v1/utilities/car/models"

	"gorm.io/gorm"
)

type ICarRepository interface {
	FindAll(limit int, skip int) ([]model.Car, *int64, error)
	FindByID(ID int) (model.Car, error)
	Create(car model.Car) (model.Car, error)
	Update(car model.Car) (model.Car, error)
	Delete(car model.Car) (model.Car, error)
}

type repository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *repository {
	return &repository{db}
}

/* func NewCustomerRepository(db *gorm.DB) *repository {
	return &repository{db}
} */

func (r *repository) FindAll(limit int, skip int) ([]model.Car, *int64, error) {
	var car []model.Car
	var count int64

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, nil, err
	}

	if err := tx.Model(&car).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := r.db.Offset(skip).Limit(limit).Order(" id asc ").Find(&car).Error; err != nil {
		return nil, nil, err
	}
	return car, &count, tx.Commit().Error
}

func (r *repository) FindByID(ID int) (model.Car, error) {
	var car model.Car
	err := r.db.Find(&car, ID).Error

	return car, err
}

func (r *repository) Create(car model.Car) (model.Car, error) {
	err := r.db.Create(&car).Error
	return car, err
}

func (r *repository) Update(car model.Car) (model.Car, error) {
	err := r.db.Save(&car).Error
	return car, err
}

func (r *repository) Delete(car model.Car) (model.Car, error) {
	err := r.db.Delete(&car).Error
	return car, err
}
