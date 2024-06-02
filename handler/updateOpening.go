package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardodf95/goopportunities/schemas"
)

// @BasePath /api/v1
// @Summary Update Opening
// @Description Update a job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Param request body UpdateOpeningRequest true "Update Opening Request"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, "ID is required")
		return
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err)
		SendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	SendSuccess(ctx, "Update Opening", opening, http.StatusOK)
}
