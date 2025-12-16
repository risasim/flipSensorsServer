package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var router *gin.Engine = gin.Default()
	SetupRoutes(router, db)

	router.Run(":8080")
}
