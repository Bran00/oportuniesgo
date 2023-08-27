package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Bran00/oportuniesgo/schemas"
)

// @BasePath /api/v1

//@Summary List Openings
//@Description List all job Openings
//@Tags Openings
//@Accept json
//@Produce json
//@Sucess 200 {object} ListOpeningResponse
//@Failure 500 {object} ErrorResponse
//@Router /openings [get]
func ShowListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(ctx , "list-opening", openings)
}
