package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marksonw/go-vacancies/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listening openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
