package driver

import (
	handler "rest-api/modules/v1/utilities/drivers/handler"
	repo "rest-api/modules/v1/utilities/drivers/repository"
	service "rest-api/modules/v1/utilities/drivers/services"
)

var (
	DriverRepository = repo.NewDriverRepository(DB)
	DriverService    = service.NewDriverService(DriverRepository)
	DriverHandler    = handler.NewDriverHandler(DriverService)
)
