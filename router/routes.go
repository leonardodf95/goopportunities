package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/leonardodf95/goopportunities/docs"
	"github.com/leonardodf95/goopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	// API v1
	v1 := router.Group(basePath)

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

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
