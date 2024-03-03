package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initilize Route
	router := gin.Default()

	// Initialize Routes
	initilizeRoutes(router)

	// Run the server
	router.Run(":3000")
}
