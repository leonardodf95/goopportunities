package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
