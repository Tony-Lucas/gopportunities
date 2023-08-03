package router

import (
	"github.com/Tony-Lucas/gopportunities/handler"
	"github.com/Tony-Lucas/gopportunities/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/")

	{
		v1.GET("/:id", handler.GetUser)
		v1.POST("/", handler.CreateUser)
		v1.PUT("/", middleware.AuthMiddleware(), handler.UpdateUser)
	}
}
