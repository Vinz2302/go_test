package routes

import (
	"rest-api/app/middlewares"
	driver "rest-api/modules/v1/utilities/drivers/handler"

	"github.com/gin-gonic/gin"
)

func Driver(router *gin.Engine, driverHandler driver.DriverHandler) {
	auth := middlewares.AuthJwt()
	role_admin := middlewares.RoleAuth(roleAdmin)
	//role_superadmin := middlewares.RoleAuth(roleSuperadmin)

	v1 := router.Group("v1/driver")

	v1.GET("", driverHandler.Index)
	v1.GET("/:id", driverHandler.FindByID)
	v1.POST("", auth, role_admin, driverHandler.Create)
	v1.PUT("/:id", auth, role_admin, driverHandler.Update)
	v1.DELETE("/:id", auth, role_admin, driverHandler.Delete)

}
