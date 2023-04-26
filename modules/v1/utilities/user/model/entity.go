package model

type Membership struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	Name          string
	DailyDiscount float32 `json:"daily_discount"`
}
