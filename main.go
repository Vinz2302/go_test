package main

import (
	"log"

	driver "rest-api/driver"
	"rest-api/modules/v1/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	/* err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	} */

	conf := driver.Conf
	if driver.ErrConf != nil {
		log.Fatal(driver.ErrConf)
	}

	router := gin.Default()

	routes.Customer(router, *driver.CustomerHandler)
	routes.Car(router, *driver.CarHandler)
	routes.Booking(router, *driver.BookingHandler)
	routes.Driver(router, *driver.DriverHandler)
	routes.Report(router, *driver.ReportHandler)
	//router.RedirectTrailingSlash = false
	//errors.Init(router)

	port := &conf.App.Port

	router.Run(":" + *port)

	//models.ConnectDatabase()

	//router.GET("/customers", controllers.GetCustomer)

	//router.Run("localhost:8080")
}
