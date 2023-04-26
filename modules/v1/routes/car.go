package routes

import (
	car "rest-api/modules/v1/utilities/car/handler"

	"github.com/gin-gonic/gin"
)

func Car(router *gin.Engine, carHandler car.CarHandler) {
	v1 := router.Group("v1/car")

	v1.GET("", carHandler.Index)
	v1.GET("/:id", carHandler.FindByID)
	v1.POST("", carHandler.Create)
	v1.PUT("/:id", carHandler.Update)
	v1.DELETE("/:id", carHandler.Delete)
}
