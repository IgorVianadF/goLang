package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDrawingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Drawing saved successfully",
	})
}
