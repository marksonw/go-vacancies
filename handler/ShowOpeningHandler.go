package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marksonw/go-vacancies/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if error := db.First(&opening, id).Error; error != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
	}

	sendSuccess(ctx, "show-opening", opening)
}
