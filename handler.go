package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func (h *Handler) dataEntry(ctx *gin.Context) {
	panic("implement me")
}

func (h *Handler) getData(ctx *gin.Context) {
	panic("implement me")
}
