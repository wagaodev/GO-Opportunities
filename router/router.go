package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initilize Route
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	// Initialize Routes
	initilizeRoutes(router)

	// Run the server
	router.Run(":3000")
}
