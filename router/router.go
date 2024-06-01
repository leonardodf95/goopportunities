package router

import "github.com/gin-gonic/gin"

func Init() {
	//initialize the router
	router := gin.Default()
	
	//initialize the routes
	initializeRoutes(router)

	router.Run(":8080") // listen and serve on
}