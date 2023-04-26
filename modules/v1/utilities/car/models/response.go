package models

type CarResponse struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	CarName        string `json:"car_name" gorm:"type:varchar(256)"`
	RentDailyPrice uint   `json:"rent_daily_price" gorm:"type:bigint"`
	Stock          uint   `json:"stock" gorm:"type:int"`
	UpdatedAt      uint   `json:"updated_at"`
	CreatedAt      uint   `json:"created_at"`
}
