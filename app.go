package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := &Handler{db: db}

	authorized := router.Group("/arduino")
	authorized.Use(AuthMiddleware())
	authorized.POST("/dataentry", h.dataEntry)

	router.GET("/dataentries", h.getData)
}
