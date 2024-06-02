package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardodf95/goopportunities/schemas"
)

// @BasePath /api/v1
// @Summary Show Opening
// @Description Show a job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, "ID is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	SendSuccess(ctx, "Show Opening", opening, http.StatusOK)
}
