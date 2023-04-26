package routes

import (
	customer "rest-api/modules/v1/utilities/customer/handler"

	"github.com/gin-gonic/gin"
)

func Customer(router *gin.Engine, customerHandler customer.CustomerHandler) {
	v1 := router.Group("v1/customer")

	v1.GET("", customerHandler.Index)
	v1.GET("/:id", customerHandler.FindById)
	v1.POST("", customerHandler.Create)
	v1.PUT("/:id", customerHandler.Edit)
	v1.DELETE("/:id", customerHandler.Delete)
}
