package models

type Customer struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Nik          uint   `json:"nik"`
	PhoneNumber  uint   `json:"phone_number"`
	MembershipID uint   `json:"membership_id"`
}
