package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wagaodev/Go-Opportunities/handler"
)

func initilizeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	v1.GET("/opening", handler.ShowOpeningHandler)
	v1.GET("/openings", handler.ListOpeningsHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
}
