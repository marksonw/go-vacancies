package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	initializeroutes(router)

	router.Run(":8080")
}
