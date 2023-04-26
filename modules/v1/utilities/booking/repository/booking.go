package repository

import (
	"fmt"
	model "rest-api/modules/v1/utilities/booking/models"
	carModel "rest-api/modules/v1/utilities/car/models"
	driver "rest-api/modules/v1/utilities/drivers/models"
	driverRepo "rest-api/modules/v1/utilities/drivers/repository"
	helper "rest-api/pkg/helpers"

	"gorm.io/gorm"
)

type IBookingRepository interface {
	FindAll(limit int, skip int) ([]model.Booking, *int64, error)
	FindByID(ID int) (model.Booking, error)
	FindByDriverID(driverID int) ([]model.Booking, error)
	Create(booking model.Booking) (*model.Booking, error)
	Update(booking model.Booking, oldBooking model.Booking) (*model.Booking, error)
	Delete(booking model.Booking) (*model.Booking, error)
	Finish(booking model.Booking) (*model.Booking, error)
}

type repository struct {
	db         *gorm.DB
	driverRepo driverRepo.IDriverRepository
}

func NewBookingRepository(db *gorm.DB, driverRepo driverRepo.IDriverRepository) *repository {
	return &repository{db, driverRepo}
}

func (r *repository) FindAll(limit int, skip int) ([]model.Booking, *int64, error) {
	var booking []model.Booking
	var count int64

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, nil, err
	}

	if err := tx.Model(&booking).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := r.db.Preload("Customer").Preload("Car").Preload("Driver").Preload("Booktype").Offset(skip).Limit(limit).Order(" id asc ").Find(&booking).Error; err != nil {
		return nil, nil, err
	}

	//fmt.Println(booking)

	return booking, &count, tx.Commit().Error
}

func (r *repository) FindByID(ID int) (model.Booking, error) {
	var booking model.Booking

	err := r.db.Preload("Customer").Preload("Car").Preload("Driver").Preload("Booktype").Find(&booking, ID).Error

	//log.Print(booking)
	return booking, err
}

func (r *repository) FindByDriverID(driversID int) ([]model.Booking, error) {
	var booking []model.Booking

	err := r.db.Model(&model.Booking{}).Where("drivers_id = ?", driversID).Scan(&booking).Error
	/* if err := nil {
		return nil, err
	} */
	/* ids := make([]int, len(bookingIDs))
	for i, bid := range bookingIDs {
		ids[i] = int(bid.ID)
	} */

	return booking, err
}

func (r *repository) Create(booking model.Booking) (*model.Booking, error) {

	var (
		car             carModel.Car
		newBooking      model.Booking
		driverIncentive driver.DriverIncentive
	)

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Find(&car, booking.CarsID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	fmt.Print("sebelum", car.Stock)

	car.Stock = car.Stock - 1

	fmt.Print("setelah", car.Stock)

	if err := tx.Save(&car).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if booking.BooktypeID == model.Driver {
		// id desc / -id
		if err := tx.Order("id DESC").Limit(1).Find(&newBooking).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		totalDriverIncentive := helper.DriverIncentive(int(booking.TotalCost))

		driverIncentive.Incentive = totalDriverIncentive

		driverIncentive.BookingID = newBooking.ID

		if err := tx.Create(&driverIncentive).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	//err := r.db.Create(&booking).Error
	return &booking, tx.Commit().Error
}

func (r *repository) Update(booking model.Booking, oldBooking model.Booking) (*model.Booking, error) {
	var (
		prevCar         carModel.Car
		car             carModel.Car
		newBooking      model.Booking
		driverIncentive driver.DriverIncentive
	)

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	//fmt.Println("booking1 = ", &booking)
	//fmt.Println("booking 2 = ", oldBooking)
	if err := tx.Find(&prevCar, oldBooking.CarsID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	fmt.Println("prev car = ", prevCar)

	prevCar.Stock = prevCar.Stock + 1

	if err := tx.Save(&prevCar).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Find(&car, booking.CarsID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	fmt.Println("car = ", car)

	car.Stock = car.Stock - 1

	if err := tx.Save(&car).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.First(&newBooking, booking.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	newBooking.CustomersID = booking.CustomersID
	newBooking.CarsID = booking.CarsID
	newBooking.StartTime = booking.StartTime
	newBooking.EndTime = booking.EndTime
	newBooking.DriversID = booking.DriversID
	newBooking.BooktypeID = booking.BooktypeID
	newBooking.TotalCost = booking.TotalCost
	newBooking.TotalDriverCost = booking.TotalDriverCost
	newBooking.Discount = booking.Discount

	if err := tx.Save(&newBooking).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//fmt.Print("booking2 = ", newBooking)

	if booking.BooktypeID == model.Driver {
		totalDriverIncentive := helper.DriverIncentive(int(booking.TotalCost))
		driverIncentive.Incentive = totalDriverIncentive
		driverIncentive.BookingID = newBooking.ID

		incentive, errIncentive := r.driverRepo.FindIncentive(int(booking.ID))

		if errIncentive != nil {
			if errIncentive == gorm.ErrRecordNotFound {
				fmt.Println("error1")
				if err := tx.Create(&driverIncentive).Error; err != nil {
					tx.Rollback()
					return nil, err
				}
				return &newBooking, tx.Commit().Error
			}
			fmt.Print("incentive error")

			return nil, errIncentive
		}

		if incentive.ID == 0 {
			fmt.Println("incentive")
		}
		incentive.Incentive = totalDriverIncentive
		incentive.BookingID = newBooking.ID

		if err := tx.Save(&incentive).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

	} else {
		fmt.Println("masuk delete")
		if err := tx.Where("booking_id = ?", booking.ID).Delete(&driverIncentive).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	/* if err := r.db.Save(&booking).Error; err != nil {
		return nil, err
	} */

	return &newBooking, tx.Commit().Error
}

func (r *repository) Delete(booking model.Booking) (*model.Booking, error) {
	var (
		driverIncentive driver.DriverIncentive
		car             carModel.Car
	)

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	fmt.Println(booking.BooktypeID)

	if booking.CarsID != 0 {
		if err := tx.Find(&car, booking.CarsID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		car.Stock = car.Stock + 1

		if err := tx.Save(&car).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if booking.BooktypeID == model.Driver {
		if err := tx.Where("booking_id = ?", booking.ID).Delete(&driverIncentive).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Delete(&booking).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//err := r.db.Delete(&booking).Error
	return &booking, tx.Commit().Error
}

func (r *repository) Finish(booking model.Booking) (*model.Booking, error) {
	var (
		car carModel.Car
	)

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	fmt.Println("after", booking)
	if err := tx.Save(&booking).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if booking.CarsID != 0 {
		if err := tx.Find(&car, booking.CarsID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		fmt.Println("car = ", car)
		car.Stock = car.Stock + 1

		if err := tx.Save(&car).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Model(&model.Booking{}).Where("id = ?", booking.ID).Update("finished", true).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &booking, tx.Commit().Error
}
