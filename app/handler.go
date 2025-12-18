package app

import (
	"database/sql"
	"flipSensorServer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

// dataEntry handles new data received from the arduino
func (h *Handler) dataEntry(ctx *gin.Context) {
	var entry database.DataEntry
	if err := ctx.ShouldBindJSON(&entry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	res, err := database.InsertEntry(h.db, entry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) getData(ctx *gin.Context) {
	res, err := database.GetEntries(h.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
