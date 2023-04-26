package models

import "time"

type CarRequest struct {
	ID             uint       `json:"id" gorm:"primaryKey" binding:"required"`
	CarName        string     `json:"car_name" gorm:"type:varchar(256)" binding:"required"`
	RentDailyPrice uint       `json:"rent_daily_price" gorm:"type:bigint" binding:"required"`
	Stock          *uint      `json:"stock" gorm:"type:int" binding:"required"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CreatedAt      *time.Time `json:"created_at"`
}
