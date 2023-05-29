package routes

import (
	"rest-api/app/middlewares"
	booking "rest-api/modules/v1/utilities/booking/handler"

	"github.com/gin-gonic/gin"
)

var (
	roleAdmin      = 1
	roleSuperadmin = 2
)

func Booking(router *gin.Engine, bookingHandler booking.BookingHandler) {

	auth := middlewares.AuthJwt()
	role_admin := middlewares.RoleAuth(roleAdmin)
	//role_superadmin := middlewares.RoleAuth(roleSuperadmin)

	v1 := router.Group("v1/booking")

	v1.GET("", auth, bookingHandler.Index)
	v1.GET("/:id", auth, bookingHandler.FindByID)
	v1.POST("", auth, role_admin, bookingHandler.Create)
	v1.PUT("/:id", auth, role_admin, bookingHandler.Update)
	v1.DELETE("/:id", auth, role_admin, bookingHandler.Delete)
	v1.PUT("/:id/finish", auth, role_admin, bookingHandler.Finish)
}
