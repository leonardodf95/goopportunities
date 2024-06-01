package main

import (
	"github.com/leonardodf95/goopportunities/config"
	"github.com/leonardodf95/goopportunities/router"
)

var (
	logger *config.Logger
)

func main(){
	logger = config.GetLogger("main")
	//Initialize configuration
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configuration %v",err)
		return
	}


	//initialize the router
	router.Init()
}