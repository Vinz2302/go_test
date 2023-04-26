package routes

import (
	booking "rest-api/modules/v1/utilities/booking/handler"

	"github.com/gin-gonic/gin"
)

func Booking(router *gin.Engine, bookingHandler booking.BookingHandler) {
	v1 := router.Group("v1/booking")

	v1.GET("", bookingHandler.Index)
	v1.GET("/:id", bookingHandler.FindByID)
	v1.POST("", bookingHandler.Create)
	v1.PUT("/:id", bookingHandler.Update)
	v1.DELETE("/:id", bookingHandler.Delete)
	v1.PUT("/:id/finish", bookingHandler.Finish)
}
