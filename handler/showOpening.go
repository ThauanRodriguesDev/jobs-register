package handler

import (
	"fmt"
	"net/http"

	"github.com/ThauanRodriguesDev/jobs-register/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show Opening
// @Description Show a Job Opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	sendSucess(ctx, "show-opening", opening)
}
