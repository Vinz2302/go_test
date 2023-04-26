package driver

import (
	customer "rest-api/modules/v1/utilities/customer/handler"
	repo "rest-api/modules/v1/utilities/customer/repository"
	service "rest-api/modules/v1/utilities/customer/services"
)

var (
	CustomerRepository = repo.NewCustomerRepository(DB)
	CustomerService    = service.NewCustomerService(CustomerRepository)
	CustomerHandler    = customer.NewCustomerHandler(CustomerService)
)
