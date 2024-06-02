package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardodf95/goopportunities/schemas"
)

// @BasePath /api/v1
// @Summary Delete Opening
// @Description Delete a job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
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

	if err := db.Delete(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "Error deleting opening")
		return
	}

	SendSuccess(ctx, "Delete Opening", opening, http.StatusOK)
}
