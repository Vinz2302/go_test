package routes

import (
	middlewares "rest-api/app/middlewares"
	customer "rest-api/modules/v1/utilities/customer/handler"

	"github.com/gin-gonic/gin"
)

func Customer(router *gin.Engine, customerHandler customer.CustomerHandler) {
	auth := middlewares.AuthJwt()
	//role_admin := middlewares.RoleAuth(roleAdmin)
	//role_superadmin := middlewares.RoleAuth(roleSuperadmin)

	v1 := router.Group("v1/customer")

	v1.GET("", auth, customerHandler.Index)
	v1.GET("/:id", auth, customerHandler.FindById)
	v1.POST("", auth, customerHandler.Create)
	v1.PUT("/:id", auth, customerHandler.Edit)
	v1.DELETE("/:id", auth, customerHandler.Delete)
}
