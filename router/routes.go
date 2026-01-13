package router

import (
	handler "github.com/IgorVianadF/goLang/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/drawing", handler.GetDrawingHandler)
		v1.POST("/drawing", handler.SaveDrawingHandler)
		v1.DELETE("/drawing", handler.DeleteDrawingHandler)
		v1.PUT("/drawing", handler.UpdateDrawingHandler)
		v1.GET("/drawings", handler.ListDrawingHandler)

	}
}
