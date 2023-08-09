package router

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,DELETE,GET,PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func Initialize() {

	// Router Instance
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.MaxMultipartMemory = 8 << 20

	// Routes

	initializeRoutes(r)
	initializeAuthRoutes(r)
	InitializeProductRoutes(r)
	InitializeAdminRoutes(r)

	// Run the server
	r.Run()
}
