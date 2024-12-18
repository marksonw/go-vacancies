package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marksonw/go-vacancies/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
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
