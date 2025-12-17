package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

// dataEntry handles new data received from the arduino
func (h *Handler) dataEntry(ctx *gin.Context) {
	var entry DataEntry
	if err := ctx.ShouldBindJSON(&entry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	res, err := insertEntry(h.db, entry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) getData(ctx *gin.Context) {
	panic("implement me")
}
