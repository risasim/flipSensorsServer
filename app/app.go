package app

import (
	"flipSensorServer/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authorized := router.Group("/arduino")
	authorized.Use(auth.AuthMiddleware())
	authorized.POST("/dataentry", dataEntry)

	router.GET("/dataentries", getData)
}

func dataEntry(ctx *gin.Context) {
	panic("implement me")
}

func getData(ctx *gin.Context) {
	panic("implement me")
}
