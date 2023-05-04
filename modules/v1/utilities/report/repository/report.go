package repository

import (
	"fmt"
	bookingModel "rest-api/modules/v1/utilities/booking/models"
	model "rest-api/modules/v1/utilities/report/models"

	"gorm.io/gorm"
)

type IReportRepository interface {
	FindMonthlyCompanyIncome(year int, month int) ([]model.Report, error)
	FindBookingActivity(year int, month int) ([]model.ReportCar, error)
	FindDriverActivity(year int, month int) ([]model.ReportDriver, error)
}

type repository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMonthlyCompanyIncome(year int, month int) ([]model.Report, error) {
	var (
		report  []model.Report
		booking bookingModel.Booking
	)

	err := r.db.Model(&booking).Select("driver_incentives.booking_id as booking_id, SUM(total_cost) as total_cost, sum(discount) as discount, sum(total_driver_cost) as total_driver_cost, sum(driver_incentives.incentive) as total_driver_incentive").
		Joins("left join driver_incentives on booking_id = bookings.id").
		Where("extract(year from bookings.end_time) = ? and extract(month from bookings.end_time) = ? and bookings.finished = true", year, month).
		Group("bookings.id, driver_incentives.booking_id").
		Scan(&report).Error

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("report = ", report)

	return report, err

}

func (r *repository) FindBookingActivity(year int, month int) ([]model.ReportCar, error) {
	var (
		booking    bookingModel.Booking
		reportCars []model.ReportCar
	)

	err := r.db.Model(&booking).
		Select("cars.id as car_id, cars.car_name, sum(date_part('day', age(bookings.end_time, bookings.start_time)) + 1) as total_days_booking, count(bookings.cars_id) as total_booking_count").
		Joins("JOIN cars ON cars.id = bookings.cars_id").
		Where("extract(year from bookings.end_time) = ? AND extract(month from bookings.end_time) = ? AND bookings.finished = true", year, month).
		Group("bookings.cars_id, cars.id").
		Scan(&reportCars).Error

	if err != nil {
		return nil, err
	}

	fmt.Println("reportCar", reportCars)
	return reportCars, err
}

func (r *repository) FindDriverActivity(year int, month int) ([]model.ReportDriver, error) {
	var (
		//booking       bookingModel.Booking
		reportDrivers []model.ReportDriver
	)

	err := r.db.Table("bookings").Select("drivers.id as driver_id, drivers.driver_name as driver_name, sum(date_part('day', age(bookings.end_time, bookings.start_time)) + 1) as total_driver_day, count(bookings.drivers_id) as total_driver_count, sum(driver_incentives.incentive) as total_incentive").
		Joins("join drivers on bookings.drivers_id = drivers.id").Joins("join driver_incentives on driver_incentives.booking_id = bookings.id").
		Where("extract(year from bookings.end_time) = ? and extract(month from bookings.end_time) = ? and bookings.finished = true", year, month).
		Group("bookings.drivers_id, drivers.id").Scan(&reportDrivers).Error

	if err != nil {
		return nil, err
	}

	fmt.Println("reportDriver", reportDrivers)
	return reportDrivers, err
}

/* func (r *repository) FindMonthlyCompanyIncome(year int, month int) (*model.Report, error) {
	var (
		report          model.Report
		booking         bookingModel.Booking
		totalDriverCost uint
		discount        float32
		totalCost       uint
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

	if err := tx.Model(&booking).Select("sum(total_cost)").Where("extract(year from end_time) = ? and extract(month from end_time) = ? and finished = true", year, month).Scan(&totalCost).Error; err != nil {
		tx.Rollback()
		fmt.Println("err", err)
		return nil, err
	}

	if err := tx.Model(&booking).Select("sum(discount)").Where("extract(year from end_time) = ? and extract(month from end_time) = ? and finished = true", year, month).Scan(&discount).Error; err != nil {
		tx.Rollback()
		fmt.Println("err", err)
		return nil, err
	}

	if err := tx.Model(&booking).Select("SUM(total_driver_cost)").Where("EXTRACT(year FROM end_time) = ? AND EXTRACT(month FROM end_time) = ? AND finished = true", year, month).Scan(&totalDriverCost).Error; err != nil {
		tx.Rollback()
		fmt.Println("err", err)
		return nil, err
	}

	rows, err := tx.Model(&booking).Select("driver_incentives.booking_id, sum(driver_incentives.incentive)").Joins("join driver_incentives on booking_id = bookings.id").Where("extract(year from bookings.end_time) = ? and extract(month from bookings.end_time) = ? and bookings.finished = true", year, month).Group("bookings.id, driver_incentives.booking_id").Rows()

	if err != nil {
		tx.Rollback()
		fmt.Println("err", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			bookingID            uint
			totalDriverIncentive uint
		)
		if err := rows.Scan(&bookingID, &totalDriverIncentive); err != nil {
			tx.Rollback()
			return nil, err
		}
		report.TotalDriverIncentive = &totalDriverIncentive
	}

	report.TotalDriverCost = &totalDriverCost
	report.Discount = &discount
	report.TotalCost = &totalCost
	fmt.Println("report", report)

	return &report, tx.Commit().Error
} */

/* func (r *repository) FindBookingActivity(year int, month int) ([]model.ReportCar, error) {
	var (
		booking    bookingModel.Booking
		reportCars []model.ReportCar
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

	rows, err := tx.Model(&booking).Select("cars.id, cars.car_name, sum(date_part('day', age(bookings.end_time, bookings.start_time))), count(bookings.cars_id)").Joins("join cars on cars.id = bookings.cars_id").Where("extract(year from bookings.end_time) = ? and extract(month from bookings.end_time) = ? and bookings.finished = true", year, month).Group("bookings.cars_id, cars.id").Rows()
	if err != nil {
		tx.Rollback()
		fmt.Println("err", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id           int
			name         string
			bookingCount int
			daysCount    int
		)
		if err := rows.Scan(&id, &name, &daysCount, &bookingCount); err != nil {
			tx.Rollback()
			return nil, err
		}

		reportCar := model.ReportCar{
			CarID:             &id,
			CarName:           &name,
			TotalBookingCount: &bookingCount,
			TotalDaysBooking:  &daysCount,
		}
		reportCars = append(reportCars, reportCar)
	}

	fmt.Println("reportCar", reportCars)
	return reportCars, tx.Commit().Error
} */

/* func (r *repository) FindDriverActivity(year int, month int) ([]model.ReportDriver, error) {
	var (
		booking       bookingModel.Booking
		reportDrivers []model.ReportDriver
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

	rows, err := tx.Model(&booking).Select("drivers.id, drivers.driver_name, sum(date_part('day', age(bookings.end_time, bookings.start_time)) + 1), count(bookings.drivers_id), sum(driver_incentives.incentive)").Joins("join drivers on bookings.drivers_id = drivers.id").Joins("join driver_incentives on driver_incentives.booking_id = bookings.id").Where("extract(year from bookings.end_time) = ? and extract(month from bookings.end_time) = ? and bookings.finished = true", year, month).Group("bookings.drivers_id, drivers.id").Rows()
	if err != nil {
		tx.Rollback()
		fmt.Println("err", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int
			name       string
			totalDay   int
			totalCount int
			Incentive  int
		)
		if err := rows.Scan(&id, &name, &totalDay, &totalCount, &Incentive); err != nil {
			tx.Rollback()
			return nil, err
		}

		reportDriver := model.ReportDriver{
			DriverID:         &id,
			DriverName:       &name,
			TotalDriverDay:   &totalDay,
			TotalDriverCount: &totalCount,
			TotalIncentive:   &Incentive,
		}
		reportDrivers = append(reportDrivers, reportDriver)
	}

	fmt.Println("reportDriver", reportDrivers)
	return reportDrivers, tx.Commit().Error
} */
