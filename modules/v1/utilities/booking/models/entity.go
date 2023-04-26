package models

import (
	"encoding/binary"
	"fmt"
	car "rest-api/modules/v1/utilities/car/models"
	customer "rest-api/modules/v1/utilities/customer/models"
	"strconv"

	driverdb "database/sql/driver"
	driver "rest-api/modules/v1/utilities/drivers/models"
	"time"
)

type Booking struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CustomersID     uint      `json:"customer_id" gorm:"type:bigint"`
	CarsID          uint      `json:"car_id" gorm:"type:bigint"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	TotalCost       uint      `json:"total_cost" gorm:"type:bigint"`
	Finished        bool      `json:"finished"`
	Discount        *float32  `json:"discount"`
	BooktypeID      BookType  `json:"booking_type_id" gorm:"type:book_type"`
	DriversID       *uint     `json:"drivers_id" gorm:"type:bigint"`
	TotalDriverCost *uint     `json:"total_driver_cost" gorm:"type:bigint"`

	Customer customer.Customer `gorm:"foreignKey:CustomersID"`
	Car      car.Car           `gorm:"foreignKey:CarsID"`
	Driver   *driver.Driver    `gorm:"foreignKey:DriversID"`
	Booktype Booktype          `gorm:"foreignKey:BooktypeID"`
	//CustomerID uint              `json:"-" gorm:"foreignKey:customers_id"`
	//CarID      uint              `json:"-" gorm:"foreignKey:cars_id"`
}

type Booktype struct {
	ID   uint
	Name string `json:"booking_type"`
}

type BookType int

const (
	Driver    BookType = 2
	NonDriver BookType = 1
)

func (s *BookType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		fmt.Println("test1")
		data := binary.BigEndian.Uint32(value.([]byte))
		*s = BookType(data)
	case string:
		fmt.Println("test2")
		data, _ := strconv.Atoi(value.(string))
		*s = BookType(data)
	default:
		fmt.Println("test3")
		return fmt.Errorf("cannot sql.Scan() BookType from: %#v", v)
	}
	return nil
}

func (s BookType) Value() (driverdb.Value, error) {
	return int(s), nil
}
