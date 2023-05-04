package models

type CustomerRequest struct {
	CustomerName string `json:"customer_name" gorm:"type:varchar(256)" binding:"required"`
	Nik          uint   `json:"nik" gorm:"type:bigint" binding:"required"`
	PhoneNumber  uint   `json:"phone_number" gorm:"type:bigint" binding:"required"`
	MembershipID uint   `json:"membership_id" gorm:"type:int" binding:"required"`
}
