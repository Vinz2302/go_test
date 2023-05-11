package driver

import (
	handler "rest-api/modules/v1/utilities/user/handler"
	repo "rest-api/modules/v1/utilities/user/repository"
	service "rest-api/modules/v1/utilities/user/service"
)

var (
	UserRepository = repo.NewUserRepository(DB)
	UserService    = service.NewUserService(UserRepository)
	UserHandler    = handler.NewUserHandler(UserService)
)
