package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Bran00/oportuniesgo/schemas"
)

func ShowListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(ctx , "list-opening", openings)
}
