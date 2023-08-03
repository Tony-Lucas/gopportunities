package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Router Instance
	r := gin.Default()

	// Routes

	initializeRoutes(r)
	initializeAuthRoutes(r)
	InitializeProductRoutes(r)
	InitializeAdminRoutes(r)
	
	// Run the server
	r.Run()
}
