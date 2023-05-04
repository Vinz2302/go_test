package models

type DriverRequest struct {
	DriverName  string `json:"driver_name" gorm:"type:varchar(256)" binding:"required"`
	Nik         uint   `json:"nik" gorm:"type:bigint" binding:"required"`
	PhoneNumber uint   `json:"phone_number" gorm:"type:bigint" binding:"required"`
	DailyCost   uint   `json:"daily_cost" gorm:"type:bigint" binding:"required"`
}
