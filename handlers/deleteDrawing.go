package handler

import (
	"fmt"
	"net/http"

	"github.com/IgorVianadF/goLang/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteDrawingHandler(ctx *gin.Context) {

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
	}

	drawing := schemas.Drawing{}

	if err := db.Where("id = ?", id).First(&drawing).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Drawing with id %s not found", id))
		return
	}

	if err := db.Delete(&drawing).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting drawing with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-drawing", drawing)
}
