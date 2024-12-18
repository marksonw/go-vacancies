package main

import (
	"github.com/marksonw/go-vacancies/config"
	"github.com/marksonw/go-vacancies/router"
)

var (
	logger *config.Logger
)

// @title           Job Openings API
// @version         1.0
// @description     This is a job openings server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
