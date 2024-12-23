package main

import (
	"github.com/ThauanRodriguesDev/jobs-register/config"
	"github.com/ThauanRodriguesDev/jobs-register/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	//Initialize Configs

	err := config.Init()

	if err != nil {
		logger.ErrorF("Config Initializate Error: %v\n", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
