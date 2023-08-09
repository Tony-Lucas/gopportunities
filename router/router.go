package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Router Instance
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://10.0.0.200:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.MaxMultipartMemory = 8 << 20

	// Routes

	initializeRoutes(r)
	initializeAuthRoutes(r)
	InitializeProductRoutes(r)
	InitializeAdminRoutes(r)

	// Run the server
	r.Run()
}
