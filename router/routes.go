package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marksonw/go-vacancies/handler"
)

func initializeroutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.POST("/opening", handler.PostOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
