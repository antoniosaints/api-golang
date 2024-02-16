package routes

import (
	"github.com/gin-gonic/gin"
)

func RoutesLojas(route *gin.Engine) *gin.RouterGroup {
	v1 := route.Group("/api/v1")
	{
		v1.GET("/lojas", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})
	}

	return v1
}
