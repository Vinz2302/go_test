package services

import (
	"errors"
	"fmt"
	model "rest-api/modules/v1/utilities/booking/models"
	repo "rest-api/modules/v1/utilities/booking/repository"
	modelCar "rest-api/modules/v1/utilities/car/models"
	carService "rest-api/modules/v1/utilities/car/services"
	modelCustomer "rest-api/modules/v1/utilities/customer/models"
	customerService "rest-api/modules/v1/utilities/customer/services"
	modelDriver "rest-api/modules/v1/utilities/drivers/models"
	driverService "rest-api/modules/v1/utilities/drivers/services"
	userService "rest-api/modules/v1/utilities/user/service"
	helper "rest-api/pkg/helpers"
	"time"
)

type IBookingService interface {
	FindAll(page int, limit int) ([]model.Booking, *int64, error)
	FindByID(id int) (model.Booking, error)
	Create(bookingRequest model.BookingRequest) (*model.Booking, error, int)
	Update(id int, BookingRequest model.BookingRequest) (*model.Booking, error, int)
	Delete(id int) (error, int)
	Finish(id int, FinishRequest model.FinishRequest) (*model.Booking, error, int)
}

type bookingService struct {
	repository      repo.IBookingRepository
	driverService   driverService.IDriverService
	carService      carService.ICarService
	customerService customerService.ICustomerService
	userService     userService.IUserService
}

func NewBookingService(repository repo.IBookingRepository, driverService driverService.IDriverService, carService carService.ICarService, userService userService.IUserService, customerService customerService.ICustomerService) *bookingService {
	return &bookingService{repository, driverService, carService, customerService, userService}
}

func (s *bookingService) FindAll(page int, limit int) ([]model.Booking, *int64, error) {
	pageA := page - 1
	if page == 0 {
		pageA = 0
	}
	skip := pageA * limit
	booking, count, err := s.repository.FindAll(limit, skip)

	return booking, count, err
}

func (s *bookingService) FindByID(id int) (model.Booking, error) {
	booking, err := s.repository.FindByID(id)
	return booking, err
}

func (s *bookingService) Create(bookingRequest model.BookingRequest) (*model.Booking, error, int) {
	var (
		totalDriverCost *uint
		discount        *float32
		err             error
		//currentRange    []time.Time
		//allRange        []time.Time
	)

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	startTime := today
	//startTime := bookingRequest.StartTime
	endTime := bookingRequest.EndTime.AddDate(0, 0, 1)

	/* for d := startTime; d.Before(endTime); d = d.AddDate(0, 0, 1) {
		currentRange = append(currentRange, d)
	} */

	days := helper.Days(endTime, startTime)
	//days := uint(endTime.Sub(startTime).Hours() / 24)

	//fmt.Print("days = ", days)

	/* customerData, errCustomer := s.customerService.FindByID(int(bookingRequest.CustomersID))
	if errCustomer != nil {
		return nil, errCustomer, 500
	}
	if customerData.ID == 0 {
		return nil, errors.New("Customer"), 404
	} */

	customerData, errCustomer, statusCode := findCustomer(s.customerService, &bookingRequest)
	if errCustomer != nil {
		return nil, errCustomer, statusCode
	}

	/* carData, errCar := s.carService.FindByID(int(bookingRequest.CarsID))
	if errCar != nil {
		return nil, errCar, 500
	}
	if carData.ID == 0 {
		return nil, errors.New("Car not exists"), 404
	}
	if carData.Stock == 0 {
		return nil, errors.New("Car out of stock"), 400
	} */
	carData, errCar, statusCode := findCar(s.carService, &bookingRequest)
	if errCar != nil {
		return nil, errCar, statusCode
	}
	fmt.Println("car data", carData)

	carDailyCost := carData.RentDailyPrice
	totalCost := helper.TotalCost(days, carDailyCost)

	membershipData, errMembership := s.userService.FindByID(int(customerData.MembershipID))
	if errMembership != nil {
		return nil, errMembership, 500
	}
	if membershipData.ID != 0 {
		discountValue := membershipData.DailyDiscount
		//discountTemp := (float32(totalCost) * discountValue)
		discountTemp := helper.Discount(float32(totalCost), discountValue)
		discount = &discountTemp
	}

	allBooking, errGetAll := s.repository.FindByDriverID(int(*bookingRequest.DriversID))
	fmt.Println("all booking ID = ", allBooking)
	if errGetAll != nil {
		return nil, errGetAll, 500
	}
	if allBooking != nil {
		return nil, errors.New("Drivers not available"), 400
	}

	/* for _, albooking := range allBooking {

		allStartDate := albooking.StartTime
		allEndDate := albooking.EndTime.AddDate(0, 0, 1)
		for d := allStartDate; d.Before(allEndDate); d = d.AddDate(0, 0, 1) {
			allRange = append(allRange, d)
		}
		fmt.Println("all start = ", allStartDate)
		fmt.Println("all end = ", allEndDate)
	} */

	/* fmt.Println("all range = ", allRange)
	for _, r := range currentRange {
		found := false
		for _, ar := range allRange {
			if r.Equal(ar) {
				found = true
				fmt.Println("overlap")
				return nil, errors.New("Driver's not available"), 400
			}
		}
		if !found {
			fmt.Println("no overlap")
		}
	} */

	if bookingRequest.BooktypeID >= model.NonDriver || bookingRequest.BooktypeID <= model.Driver {
		if bookingRequest.BooktypeID == model.Driver {
			/* driverData, errDriver := s.driverService.FindByID(int(*bookingRequest.DriversID))
			if errDriver != nil {
				return nil, errDriver, 500
			}
			if driverData.ID == 0 {
				return nil, errors.New("Driver"), 404
			}
			dailyCostDriver := driverData.DailyCost

			fmt.Print("daily cost = ", dailyCostDriver)

			//totalDriverCostTemp := (days * dailyCostDriver)
			totalDriverCostTemp := helper.TotalDriverCost(days, dailyCostDriver)
			totalDriverCost = &totalDriverCostTemp */
			totalDriverCost, err = calculateDriverCost(s.driverService, &bookingRequest, days)
			if err != nil {
				return nil, err, 500
			}
		}
		if bookingRequest.BooktypeID == model.NonDriver {
			bookingRequest.DriversID = nil
		}
	} else {
		return nil, errors.New("Booking Type"), 404
	}

	fmt.Print("car cost = ", carDailyCost)

	//fmt.Print("driver daily cost = ", driverDailyCost)
	//fmt.Print("Total Cost = ", totalCost)
	//fmt.Printf("Total Driver Cost = %v\n", *totalDriverCost)
	//fmt.Printf("discount = %v\n", *discount)

	Booking := model.Booking{
		CustomersID:     bookingRequest.CustomersID,
		CarsID:          bookingRequest.CarsID,
		StartTime:       startTime,
		EndTime:         bookingRequest.EndTime,
		BooktypeID:      bookingRequest.BooktypeID,
		DriversID:       bookingRequest.DriversID,
		TotalCost:       totalCost,
		TotalDriverCost: totalDriverCost,
		Discount:        discount,
	}

	newBooking, err := s.repository.Create(Booking)
	if err != nil {
		return nil, err, 500
	}

	return newBooking, nil, 200
}

func (s *bookingService) Update(id int, BookingRequest model.BookingRequest) (*model.Booking, error, int) {
	var (
		discount        *float32
		totalDriverCost *uint
		oldBooking      model.Booking
		err             error
		//currentRange    []time.Time
		//allRange        []time.Time
		//bookedRange     []time.Time
		//newDriver uint
		//oldDriver uint
	)

	booking, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return nil, errGet, 500
	}
	if booking.ID == 0 {
		return nil, errors.New("ID"), 404
	}
	if booking.Finished == true {
		return nil, errors.New("Booking already finished"), 400
	}

	//startTime := BookingRequest.StartTime
	startTime := booking.StartTime
	endTime := BookingRequest.EndTime.AddDate(0, 0, 1)

	if startTime.After(endTime) {
		return nil, errors.New("Invalid end date"), 400
	}

	//days := uint(endTime.Sub(startTime).Hours() / 24)
	days := helper.Days(endTime, startTime)

	/* for d := startTime; d.Before(endTime); d = d.AddDate(0, 0, 1) {
		currentRange = append(currentRange, d)
		fmt.Println("day", d)
	} */

	/* customerData, errCustomer := s.customerService.FindByID(int(BookingRequest.CustomersID))
	if errCustomer != nil {
		return nil, errCustomer, 500
	}
	if customerData.ID == 0 {
		return nil, errors.New("Customer"), 404
	} */
	customerData, errCustomer, statusCode := findCustomer(s.customerService, &BookingRequest)
	if errCustomer != nil {
		return nil, errCustomer, statusCode
	}

	carData, errCar, statusCode := findCar(s.carService, &BookingRequest)
	if errCar != nil {
		return nil, errCar, statusCode
	}

	carDailyCost := carData.RentDailyPrice
	totalCost := helper.TotalCost(days, carDailyCost)

	membershipData, errMembership := s.userService.FindByID(int(customerData.MembershipID))
	if errMembership != nil {
		return nil, errMembership, 500
	}
	if membershipData.ID != 0 {
		discountValue := membershipData.DailyDiscount
		discountTemp := helper.Discount(float32(totalCost), discountValue)
		discount = &discountTemp
	}

	/* bookedStart := startTime
	bookedEnd := booking.EndTime.AddDate(0, 0, 1)
	for d := bookedStart; d.Before(bookedEnd); d = d.AddDate(0, 0, 1) {
		bookedRange = append(bookedRange, d)
	} */
	//fmt.Println("current range", currentRange)
	//fmt.Println("booked range = ", bookedRange)

	if BookingRequest.BooktypeID == model.Driver {

		if booking.DriversID == nil {
			booking.DriversID = BookingRequest.DriversID
		}

		//oldDriver = *booking.DriversID
		//newDriver = *BookingRequest.DriversID
		//fmt.Println("oldDriver", oldDriver)
		//fmt.Println("newDriver", newDriver)

		/* if oldDriver != newDriver {
			fmt.Println("test !=")
			allBooking, errGetAll := s.repository.FindByDriverID(int(*BookingRequest.DriversID))
			if errGetAll != nil {
				return nil, err, 500
			}

			fmt.Println("allbooking ID= ", allBooking)

			for _, cr := range currentRange {
				for _, booking := range allBooking {
					allStartDate := booking.StartTime
					allEndDate := booking.EndTime.AddDate(0, 0, 1)
					for d := allStartDate; d.Before(allEndDate); d = d.AddDate(0, 0, 1) {
						allRange = append(allRange, d)
					}
				}
				for _, ar := range allRange {
					if cr.Equal(ar) {
						//fmt.Println("currentrange1", currentRange)
						//fmt.Println("allrange1", allRange)
						fmt.Println("overlap 2")
						return nil, errors.New("Driver not available"), 400
					}
				}
			}
		} */

		/* if oldDriver == newDriver {
			fmt.Println("test ==")
			for _, cr := range currentRange {
				found := false
				for _, br := range bookedRange {
					if cr.Equal(br) {
						found = true
						fmt.Println("overlap")
						break
					}
				}
				if !found {
					fmt.Println("no overlap")
					allBooking, errGetAll := s.repository.FindByDriverID(int(*BookingRequest.DriversID))
					if errGetAll != nil {
						return nil, err, 500
					}

					fmt.Println("allbooking ID= ", allBooking)

					for _, booking := range allBooking {

						allStartDate := booking.StartTime
						allEndDate := booking.EndTime.AddDate(0, 0, 1)
						for d := allStartDate; d.Before(allEndDate); d = d.AddDate(0, 0, 1) {
							allRange = append(allRange, d)
						}
						//fmt.Println("all start = ", allStartDate)
						//fmt.Println("all end = ", allEndDate)
					}

					//fmt.Println("current range = ", currentRange)
					//fmt.Println("all range = ", allRange)

					for _, ar := range allRange {
						if cr.Equal(ar) {
							found = true
							fmt.Println("overlap 2")
							return nil, errors.New("Driver not available"), 400
						}
					}
					fmt.Println("no overlap 2")
				}
			}
		} */

		totalDriverCost, err = calculateDriverCost(s.driverService, &BookingRequest, days)
		if err != nil {
			return nil, err, 500
		}
	}

	/* var nowRange []time.Time
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	fmt.Println("current Time = ", today)
	tomorrow, _ := time.Parse("2006-01-02 15:04:05", "2023-04-19 00:00:00")
	fmt.Println("tomorrow = ", tomorrow)
	for d := today; d.Before(tomorrow); d = d.AddDate(0, 0, 1) {
		nowRange = append(nowRange, d)
	}
	fmt.Println("now range = ", nowRange) */

	/* if BookingRequest.BooktypeID == model.Driver {
		totalDriverCost, err = calculateDriverCost(s.driverService, &BookingRequest, days)
		if err != nil {
			return nil, err, 500
		}
	} */

	/* if BookingRequest.BooktypeID == model.NonDriver {
		BookingRequest.DriversID = nil
	} */

	//fmt.Print("car cost = ", carDailyCost)

	//fmt.Print("Total Cost = ", totalCost)
	//fmt.Print("Total Driver Cost = ", *totalDriverCost)
	//fmt.Print("discount = ", *discount)

	oldBooking = booking

	booking.CustomersID = BookingRequest.CustomersID
	booking.CarsID = BookingRequest.CarsID
	booking.StartTime = startTime
	booking.EndTime = BookingRequest.EndTime
	booking.BooktypeID = BookingRequest.BooktypeID
	booking.DriversID = BookingRequest.DriversID
	booking.TotalCost = totalCost
	booking.TotalDriverCost = totalDriverCost
	booking.Discount = discount

	//fmt.Println("old booking = ", oldBooking)
	//fmt.Println("current booking = ", booking)

	newBooking, err := s.repository.Update(booking, oldBooking)
	if err != nil {
		return nil, err, 500
	}
	//fmt.Print("newBooking = ", newBooking)

	return newBooking, err, 200
}

func (s *bookingService) Delete(id int) (error, int) {

	Booking, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return errGet, 500
	}
	if Booking.ID == 0 {
		return errors.New("ID"), 404
	}
	if Booking.Finished == true {
		return errors.New("Booking already finished"), 400
	}

	_, err := s.repository.Delete(Booking)
	if err != nil {
		return err, 500
	}
	return nil, 200
}

func (s *bookingService) Finish(id int, FinishRequest model.FinishRequest) (*model.Booking, error, int) {
	var (
		discount        *float32
		totalDriverCost *uint
		endTime         time.Time
	)

	//now := time.Now().UTC()
	//today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	time := FinishRequest.EndTime

	fmt.Println("time = ", time)

	booking, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return nil, errGet, 500
	}
	if booking.ID == 0 {
		return nil, errors.New("ID"), 404
	}
	if booking.Finished == true {
		return nil, errors.New("Booking already finished"), 400
	}

	customerData, errCustomer := s.customerService.FindByID(int(booking.CustomersID))
	if errCustomer != nil {
		return nil, errCustomer, 500
	}
	if customerData.ID == 0 {
		return nil, errors.New("Customer"), 404
	}

	carData, errCar := s.carService.FindByID(int(booking.CarsID))
	if errCar != nil {
		return nil, errCar, 500
	}

	fmt.Println("endtime", booking.EndTime)
	fmt.Println("today", time)

	if booking.EndTime.Before(time) {
		endTime = time
	} else if booking.EndTime.After(time) {
		endTime = booking.EndTime
	} else {
		endTime = time
	}

	startTime := booking.StartTime
	if startTime.After(FinishRequest.EndTime) {
		return nil, errors.New("Invalid end date"), 400
	}
	days := helper.Days(endTime, startTime)
	carDailyCost := carData.RentDailyPrice
	totalCost := helper.TotalCost(days, carDailyCost)

	memberhipData, errMembership := s.userService.FindByID(int(customerData.MembershipID))
	if errMembership != nil {
		return nil, errMembership, 500
	}
	if memberhipData.ID != 0 {
		discountValue := memberhipData.DailyDiscount
		discountTemp := helper.Discount(float32(totalCost), discountValue)
		discount = &discountTemp
	}

	if booking.BooktypeID == model.Driver {
		driverData, errDriver := s.driverService.FindByID(int(*booking.DriversID))
		if errDriver != nil {
			return nil, errDriver, 500
		}
		if driverData.ID == 0 {
			return nil, errors.New("Driver not found"), 500
		}
		dailyCostDriver := driverData.DailyCost
		totalDriverCostTemp := helper.TotalDriverCost(days, dailyCostDriver)
		totalDriverCost = &totalDriverCostTemp

	}

	fmt.Println("before = ", booking)

	booking.EndTime = time
	booking.TotalCost = totalCost
	booking.TotalDriverCost = totalDriverCost
	booking.Discount = discount

	newBooking, err := s.repository.Finish(booking)
	if err != nil {
		return nil, err, 500
	}
	return newBooking, nil, 200
}

func calculateDriverCost(driverService driverService.IDriverService, bookingRequest *model.BookingRequest, days uint) (*uint, error) {
	driverData, errDriver := driverService.FindByID(int(*bookingRequest.DriversID))
	if errDriver != nil {
		return nil, errDriver
	}
	if driverData.ID == 0 {
		return nil, errors.New("Driver not found")
	}
	dailyCostDriver := driverData.DailyCost
	totalDriverCostTemp := helper.TotalDriverCost(days, dailyCostDriver)
	totalDriverCost := &totalDriverCostTemp
	return totalDriverCost, errDriver
}

func findDriver(driverService driverService.IDriverService, bookingRequest *model.BookingRequest) (*modelDriver.Driver, error, int) {
	driverData, errDriver := driverService.FindByID(int(*bookingRequest.DriversID))
	if errDriver != nil {
		return nil, errDriver, 500
	}
	if driverData.ID == 0 {
		return nil, errors.New("Driver not found"), 404
	}
	return &driverData, nil, 200
}

func findCustomer(customerService customerService.ICustomerService, bookingRequest *model.BookingRequest) (*modelCustomer.Customer, error, int) {
	customerData, errCustomer := customerService.FindByID(int(bookingRequest.CustomersID))
	if errCustomer != nil {
		return nil, errCustomer, 500
	}
	if customerData.ID == 0 {
		return nil, errors.New("ID not found"), 404
	}
	return &customerData, nil, 200
}

func findCar(carService carService.ICarService, bookingRequest *model.BookingRequest) (*modelCar.Car, error, int) {
	carData, errCar := carService.FindByID(int(bookingRequest.CarsID))
	fmt.Println("car =", carData)
	if errCar != nil {
		return nil, errCar, 500
	}
	if carData.ID == 0 {
		return nil, errors.New("Car ID not found"), 404
	}
	if uint(carData.Stock) == 0 {
		return nil, errors.New("Car out of stock"), 400
	}
	return &carData, nil, 200
}
