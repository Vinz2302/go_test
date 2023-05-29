package routes

import (
	"rest-api/app/middlewares"
	report "rest-api/modules/v1/utilities/report/handler"

	"github.com/gin-gonic/gin"
)

func Report(router *gin.Engine, reportHandler report.ReportHandler) {
	auth := middlewares.AuthJwt()
	//role_admin := middlewares.RoleAuth(roleAdmin)
	role_superadmin := middlewares.RoleAuth(roleSuperadmin)

	v1 := router.Group("v1/report")

	v1.GET("", auth, role_superadmin, reportHandler.FindMonthlyCompanyIncome)
	v1.GET("/car", auth, role_superadmin, reportHandler.FindBookingActivity)
	v1.GET("/driver", auth, role_superadmin, reportHandler.FindDriverActivity)
}
