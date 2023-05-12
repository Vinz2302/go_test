package routes

import (
	user "rest-api/modules/v1/utilities/user/handler"

	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine, userHandler user.UserHandler) {

	v1 := router.Group("v1/user")

	v1.GET("/:id", userHandler.GetById)
	v1.POST("", userHandler.Create)
	v1.POST("/login", userHandler.Login)

}
