package routes

import (
	"rest-api/app/middlewares"
	car "rest-api/modules/v1/utilities/car/handler"

	"github.com/gin-gonic/gin"
)

func Car(router *gin.Engine, carHandler car.CarHandler) {
	auth := middlewares.AuthJwt()
	role_admin := middlewares.RoleAuth(roleAdmin)
	//role_superadmin := middlewares.RoleAuth(roleSuperadmin)

	v1 := router.Group("v1/car")

	v1.GET("", carHandler.Index)
	v1.GET("/:id", carHandler.FindByID)
	v1.POST("", auth, role_admin, carHandler.Create)
	v1.PUT("/:id", auth, role_admin, carHandler.Update)
	v1.DELETE("/:id", auth, carHandler.Delete)
}
