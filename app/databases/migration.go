package databases

import (
	booking "rest-api/modules/v1/utilities/booking/models"
	car "rest-api/modules/v1/utilities/car/models"
	customer "rest-api/modules/v1/utilities/customer/models"
	driver "rest-api/modules/v1/utilities/drivers/models"

	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Print("start migrate")
	db.AutoMigrate(
		&booking.Booking{},
	)
	db.AutoMigrate(&car.Car{})
	db.AutoMigrate(&customer.Customer{})
	db.AutoMigrate(&driver.Driver{})
	db.AutoMigrate(&customer.Membership{})
	db.AutoMigrate(&booking.Booktype{})
	db.AutoMigrate(&driver.DriverIncentive{})

}
