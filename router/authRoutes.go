package router

import (
	"github.com/Tony-Lucas/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeAuthRoutes(r *gin.Engine) {
	v1 := r.Group("/auth")
	{
		v1.GET("", handler.VerifyAuth)
		v1.POST("", handler.Auth)
	}
}
