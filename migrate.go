package main

import (
	database "rest-api/app/databases"
	driver "rest-api/driver"
)

func main3() {

	db := driver.DB
	database.Migrate(db)

}
