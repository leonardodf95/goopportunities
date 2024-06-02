package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardodf95/goopportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "Error listing openings")
		return
	}

	SendSuccess(ctx, "List Openings", openings, http.StatusOK)
}
