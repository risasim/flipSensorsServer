package main

import (
	"flipSensorServer/app"
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	app.SetupRoutes(router)

	router.Run(":8080")
}
