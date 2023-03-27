package main

import (
	"Vinz2302/go_test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/customers", controllers.getCustomer)

	router.Run("localhost:8080")
}
