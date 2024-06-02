package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardodf95/goopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	// API v1
	v1 := router.Group("/api/v1")

	{
		// Show Opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		// Create Opening
		v1.POST("/opening", handler.CreateOpeningHandler)
		// Delete Opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// Update Opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		// List Openings
		v1.GET("/openings", handler.ListOpeningHandler)
	}

}
