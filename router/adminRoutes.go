package router

import (
	"github.com/Tony-Lucas/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func InitializeAdminRoutes(r *gin.Engine){

	v1 := r.Group("/admin")
	{
		v1.POST("/",handler.CreateAdmin)
	}
}	