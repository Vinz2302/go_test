package driver

import (
	"rest-api/modules/v1/utilities/user/repository"
	helperDatabases "rest-api/pkg/helpers/databases"
)

var (
	HelperDatabase = helperDatabases.InitHelperDatabase(DB)
	UserRepository = repository.InitUserRepository(DB, HelperDatabase, &Conf)
)
