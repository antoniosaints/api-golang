package routes

import (
	ctr "api/controllers"

	"github.com/gin-gonic/gin"
)

func RoutesUsers(route *gin.Engine) *gin.RouterGroup {

	v1 := route.Group("/api/v1")
	userController := ctr.UserControllers()

	{
		v1.GET("/user", userController.GetUsers)
		v1.GET("/user/:name", userController.GetUser)
	}

	return v1
}
