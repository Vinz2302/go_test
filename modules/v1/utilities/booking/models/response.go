package models

import "time"

type BookingResponse struct {
	ID              uint
	CustomersID     uint      `json:"customer_id"`
	CarsID          uint      `json:"car_id"`
	DriversID       *uint     `json:"driver_id"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	TotalCost       uint      `json:"total_cost" gorm:"type:bigint"`
	TotalDriverCost *uint     `json:"total_driver_cost" gorm:"type:bigint"`
	Finished        bool      `json:"finished"`

	Discount *float32 `json:"discount"`
	Booktype BookType `json:"booking_type"`
}
