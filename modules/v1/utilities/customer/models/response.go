package models

type CustomerResponse struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CustomerName string `json:"customer_name" gorm:"type:varchar(256)" binding:"required"`
	Nik          uint   `json:"nik" gorm:"type:bigint"`
	PhoneNumber  uint   `json:"phone_number" gorm:"type:bigint"`
	MembershipID uint   `json:"membership_id" gorm:"type:int"`
}
