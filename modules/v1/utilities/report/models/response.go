package models

type ReportResponse struct {
	BookingID            *uint `json:"booking_id"`
	TotalDriverCost      *uint `json:"total_driver_cost" gorm:"column:total_driver_cost"`
	TotalDriverIncentive *uint `json:"total_driver_incentive"`
	TotalDriverExpense   *uint `json:"total_driver_expense"`
	TotalGrossIncome     *uint `json:"total_gross_income"`
	TotalNettIncome      *uint `json:"total_nett_income"`
}
