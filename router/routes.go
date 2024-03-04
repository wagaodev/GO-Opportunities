package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wagaodev/Go-Opportunities/handler"
)

func initilizeRoutes(router *gin.Engine) {
	// InitializeHandler
	handler.InitializeHandler()
	basePath := "/api/v1"

	// Routes
	v1 := router.Group(basePath)
	v1.GET("/opening", handler.ShowOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.GET("/openings", handler.ListOpeningsHandler)
}
