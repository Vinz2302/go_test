package model

import (
	role "rest-api/modules/v1/utilities/role/model"
	"time"

	"gorm.io/gorm"
)

type UserStatusType string

/* const (
	Active UserStatusType = "active"
	Resign UserStatusType = "resign"
) */

/* type Users struct {
	ID           *uint   `gorm:"primaryKey"`
	AuthServerId *uint   `gorm:"unique; not null"`
	Nip          *string `gorm:"unique; type:varchar(256); not null"`
	Name         *string `gorm:"type:varchar(256);"`
	Email        *string `gorm:"type:varchar(256);"`
	RoleId       *uint
	Status       UserStatusType `gorm:"type:sales_zone_type"`
	CreatedAt    time.Time      `gorm:"DEFAULT:current_timestamp"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Role         role.Roles `json:"role" gorm:"foreignKey:RoleId"`
} */

type Users struct {
	ID        *uint   `gorm:"primaryKey"`
	Name      *string `gorm:"unique; not null"`
	Email     *string `gorm:"unique; type:varchar(256); not null"`
	roleId    *uint
	CreatedAt time.Time `gorm:"DEFAULT:current_timestamp"`
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
	Role      role.Roles `json:"role" gorm:"foreignKey:RoleId`
}

/* func (user *UserStatusType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*user = UserStatusType(value.([]byte))
	case string:
		*user = UserStatusType(value.(string))
	default:
		return fmt.Errorf("cannot sql.Scan() UserStatusType from: %#v", v)
	}
	return nil
} */

/* func (user UserStatusType) Value() (driver.Value, error) {
	return string(user), nil
} */
