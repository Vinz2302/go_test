package driver

import (
	handler "rest-api/modules/v1/utilities/booking/handler"
	repo "rest-api/modules/v1/utilities/booking/repository"
	service "rest-api/modules/v1/utilities/booking/services"
)

var (
	BookingRepository = repo.NewBookingRepository(DB, DriverRepository)
	BookingService    = service.NewBookingService(BookingRepository, DriverService, CarService, CustomerService, CustomerRepository)
	BookingHandler    = handler.NewBookingHandler(BookingService)
)
