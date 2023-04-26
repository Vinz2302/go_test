package driver

import (
	config "rest-api/app/config"
	database "rest-api/app/databases"
)

var (
	Conf, ErrConf = config.Init()

	DB, _ = database.Connect(&Conf)
)
