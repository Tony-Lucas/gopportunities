package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Router Instance
	r := gin.Default()
	r.Use(cors.Default())
	r.MaxMultipartMemory = 8 << 20

	// Routes

	initializeRoutes(r)
	initializeAuthRoutes(r)
	InitializeProductRoutes(r)
	InitializeAdminRoutes(r)

	// Run the server
	r.Run()
}
