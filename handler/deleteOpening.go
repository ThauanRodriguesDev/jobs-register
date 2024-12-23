package handler

import (
	"fmt"
	"net/http"

	"github.com/ThauanRodriguesDev/jobs-register/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete Opening
// @Description Delete Job
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	// Find Opening

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found\n", id))
		return
	}
	// Delete Opening

	if erro := db.Delete(&opening).Error; erro != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSucess(ctx, "delete-opening", opening)
}
