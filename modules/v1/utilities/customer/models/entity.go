package models

import (
	membership "rest-api/modules/v1/utilities/user/model"
)

/* type Customer struct {
	ID           uint       //`json:"id" gorm:"primaryKey"`
	CustomerName string     `json:"customer_name" gorm:"type:varchar(256)"`
	Nik          uint       `json:"nik" gorm:"type:bigint"`
	PhoneNumber  uint       `json:"phone_number" gorm:"type:bigint"`
	MembershipID uint       `json:"membership_id" gorm:"type:int"`
	UpdatedAt    *time.Time `json:"updated_at"`
	CreatedAt    *time.Time `json:"created_at"`
	IsDeleted    bool       `json:"is_deleted"`
	DeletedAt    *time.Time `json:"deleted_at"`
} */

type Customer struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CustomerName string `json:"customer_name" gorm:"type:varchar(256)"`
	Nik          uint   `json:"nik" gorm:"type:bigint"`
	PhoneNumber  uint   `json:"phone_number" gorm:"type:bigint"`
	MembershipID uint   `json:"membership_id" gorm:"type:int"`

	Membership membership.Membership `gorm:"foreignKey:MembershipID"`

	//Membership membership `gorm:"foreignKey:MembershipID"`
}

/* type membership struct {
	ID            uint
	Name          string
	DailyDiscount float32 `json:daily_discount`
}
*/
