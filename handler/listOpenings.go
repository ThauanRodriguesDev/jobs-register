package handler

import (
	"net/http"

	"github.com/ThauanRodriguesDev/jobs-register/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Openings
// @Description List all Jobs
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opening")
		return
	}

	sendSucess(ctx, "list-openings", openings)
}
