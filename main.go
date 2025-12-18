package main

import (
	"flipSensorServer/app"
	db2 "flipSensorServer/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db2.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var router *gin.Engine = gin.Default()
	app.SetupRoutes(router, db)

	router.Run(":8080")
}
