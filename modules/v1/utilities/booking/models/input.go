package models

import "time"

type BookingRequest struct {
	CustomersID uint      `json:"customer_id" binding:"required"`
	CarsID      uint      `json:"car_id" binding:"required"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	BooktypeID  BookType  `json:"booking_type_id" gorm:"type:book_type" binding:"required"`
	DriversID   *uint     `json:"driver_id"`
}
