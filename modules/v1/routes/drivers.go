package routes

import (
	driver "rest-api/modules/v1/utilities/drivers/handler"

	"github.com/gin-gonic/gin"
)

func Driver(router *gin.Engine, driverHandler driver.DriverHandler) {
	v1 := router.Group("v1/driver")

	v1.GET("", driverHandler.Index)
	v1.GET("/:id", driverHandler.FindByID)
	v1.POST("", driverHandler.Create)
	v1.PUT("/:id", driverHandler.Update)
	v1.DELETE("/:id", driverHandler.Delete)

}
