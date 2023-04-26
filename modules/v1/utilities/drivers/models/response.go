package models

type DriverResponse struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	DriverName  string `json:"driver_name" gorm:"type:varchar(256)"`
	Nik         uint   `json:"nik" gorm:"type:bigint"`
	PhoneNumber uint   `json:"phone_number" gorm:"type:bigint"`
	DailyCost   uint   `json:"daily_cost" gorm:"type:bigint"`
}
