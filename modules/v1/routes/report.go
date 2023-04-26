package routes

import (
	report "rest-api/modules/v1/utilities/report/handler"

	"github.com/gin-gonic/gin"
)

func Report(router *gin.Engine, reportHandler report.ReportHandler) {
	v1 := router.Group("v1/report")

	v1.GET("", reportHandler.FindMonthlyCompanyIncome)
	v1.GET("/car", reportHandler.FindBookingActivity)
	v1.GET("/driver", reportHandler.FindDriverActivity)
}
