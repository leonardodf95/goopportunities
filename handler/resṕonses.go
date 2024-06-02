package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leonardodf95/goopportunities/schemas"
)

func SendError(ctx *gin.Context, status int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"message": message,
		"status":  status,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}, status int) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"message": fmt.Sprintf("%s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
