package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wagaodev/Go-Opportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameters").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	sendSuccess(ctx, "show-openingi", opening)
}
