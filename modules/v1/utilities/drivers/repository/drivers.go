package repository

import (
	model "rest-api/modules/v1/utilities/drivers/models"

	"gorm.io/gorm"
)

type IDriverRepository interface {
	FindAll(limit int, skip int) ([]model.Driver, *int64, error)
	FindByID(ID int) (model.Driver, error)
	FindIncentive(id int) (model.DriverIncentive, error)
	Create(driver model.Driver) (model.Driver, error)
	Update(driver model.Driver) (model.Driver, error)
	Delete(driver model.Driver) (model.Driver, error)
}

type repository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindIncentive(id int) (model.DriverIncentive, error) {
	var driverIncentive model.DriverIncentive

	err := r.db.Where("booking_id = ?", id).First(&driverIncentive).Error

	return driverIncentive, err
}

func (r *repository) FindAll(limit int, skip int) ([]model.Driver, *int64, error) {
	var driver []model.Driver
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

	if err := tx.Model(&driver).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := r.db.Offset(skip).Limit(limit).Order("id asc").Find(&driver).Error; err != nil {
		return nil, nil, err
	}
	return driver, &count, tx.Commit().Error
}

func (r *repository) FindByID(ID int) (model.Driver, error) {
	var driver model.Driver
	err := r.db.Find(&driver, ID).Error

	return driver, err
}

func (r *repository) Create(driver model.Driver) (model.Driver, error) {
	err := r.db.Create(&driver).Error

	return driver, err
}

func (r *repository) Update(driver model.Driver) (model.Driver, error) {
	err := r.db.Save(&driver).Error

	return driver, err
}

func (r *repository) Delete(driver model.Driver) (model.Driver, error) {
	err := r.db.Delete(&driver).Error

	return driver, err
}
