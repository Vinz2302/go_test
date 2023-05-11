package main

import (
	"log"

	driver "rest-api/driver"
	"rest-api/modules/v1/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	conf := driver.Conf
	if driver.ErrConf != nil {
		log.Fatal(driver.ErrConf)
	}

	router := gin.Default()

	routes.Customer(router, *driver.CustomerHandler)
	routes.Car(router, *driver.CarHandler)
	routes.Booking(router, *driver.BookingHandler)
	routes.Driver(router, *driver.DriverHandler)
	routes.User(router, *driver.UserHandler)
	routes.Report(router, *driver.ReportHandler)

	port := &conf.App.Port

	router.Run(":" + *port)

}
