package driver

import (
	repo "rest-api/modules/v1/utilities/user/repository"
	service "rest-api/modules/v1/utilities/user/service"
)

var (
	UserRepository = repo.NewUserRepository(DB)
	UserService    = service.NewMembershipService(UserRepository)
)
