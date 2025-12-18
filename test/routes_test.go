package test

import (
	"database/sql"
	"flipSensorServer/app"
	"flipSensorServer/database"
	"testing"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *sql.DB) {
	gin.SetMode(gin.TestMode)

	db := database.SetupTestDB(t)
	defer db.Close()

	router := gin.New()
	app.SetupRoutes(router, db)

	return router, db
}
