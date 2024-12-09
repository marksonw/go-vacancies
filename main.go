package main

import (
	"github.com/marksonw/go-vacancies/config"
	"github.com/marksonw/go-vacancies/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Configuration
	err := config.Init()

	if err != nil {
		logger.Errorf("configuration initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
