package main

import (
	database "rest-api/app/databases"
	driver "rest-api/driver"
)

func main3() {

	db := driver.DB
	database.Migrate(db)

}

//migrate create -ext sql -dir app/databases/migrations/ -seq setName_mg
//migrate -path app/databases/migrations/ -database "postgresql://postgres:neverdie@localhost:5432/go_test1?sslmode=disable" -verbose up
//migrate -path app/databases/migrations/ -database "postgresql://postgres:neverdie@localhost:5432/go_test1?sslmode=disable" force 2
