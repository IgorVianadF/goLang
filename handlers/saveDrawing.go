package handler

import (
	"net/http"

	"github.com/IgorVianadF/goLang/schemas"
	"github.com/gin-gonic/gin"
)

func SaveDrawingHandler(ctx *gin.Context) {
	request := SaveDrawingRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Erro de validação: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	drawing := schemas.Drawing{
		Name:      request.Name,
		Price:     request.Price,
		Customer:  request.Customer,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
	}

	if err := db.Create(&drawing).Error; err != nil {
		logger.Errorf("Erro ao persistir desenho no banco: %v", err.Error())
		return
	}

	sendSuccess(ctx, "save-drawing", drawing)
}
