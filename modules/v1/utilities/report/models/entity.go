package models

type Report struct {
	BookingID            *uint    `json:"booking_id"`
	TotalDriverCost      *uint    `json:"total_driver_cost"`
	TotalDriverIncentive *uint    `json:"total_driver_incentive"`
	TotalDriverExpense   *uint    `json:"total_driver_expense"`
	TotalGrossIncome     *uint    `json:"total_gross_income"`
	TotalNettIncome      *uint    `json:"total_nett_income"`
	Discount             *float32 `json:"discount"`
	TotalCost            *uint    `json:"total_cost"`
}

type ReportCar struct {
	CarID             *int    `json:"car_id"`
	CarName           *string `json:"car_name"`
	TotalDaysBooking  *int    `json:"total_days_count"`
	TotalBookingCount *int    `json:"total_booking_count"`
}

type ReportDriver struct {
	DriverID         *int
	DriverName       *string
	TotalDriverDay   *int
	TotalDriverCount *int
	TotalIncentive   *int
}
