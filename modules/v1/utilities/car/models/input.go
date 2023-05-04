package models

type CarRequest struct {
	CarName        string `json:"car_name" gorm:"type:varchar(256)" binding:"required"`
	RentDailyPrice uint   `json:"rent_daily_price" gorm:"type:bigint" binding:"required"`
	Stock          *uint  `json:"stock" gorm:"type:int" binding:"required"`
}
