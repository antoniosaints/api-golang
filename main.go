package main

import (
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.RoutesUsers(app)
	routes.RoutesLojas(app)

	app.Run("0.0.0.0:8080")
}
