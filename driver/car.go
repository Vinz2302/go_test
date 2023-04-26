package driver

import (
	handler "rest-api/modules/v1/utilities/car/handler"
	repo "rest-api/modules/v1/utilities/car/repository"
	service "rest-api/modules/v1/utilities/car/services"
)

var (
	CarRepository = repo.NewCarRepository(DB)
	CarService    = service.NewCarService(CarRepository)
	CarHandler    = handler.NewCarHandler(CarService)
)
