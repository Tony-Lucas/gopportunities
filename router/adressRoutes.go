package router

import (
	"github.com/Tony-Lucas/gopportunities/handler"
	"github.com/Tony-Lucas/gopportunities/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeAdressRoutes(r *gin.Engine){

	v1 := r.Group("/adress")
	{
		v1.GET("/:userId",middleware.AuthMiddleware(),handler.GetUserAdresses)
		v1.GET("/:id/:userId",middleware.AuthMiddleware(),handler.GetUserAdress)
		v1.POST("/",middleware.AuthMiddleware(),handler.PostUserAdress)
		v1.PUT("/",middleware.AuthMiddleware(),handler.UpdateUserAdress)
		v1.DELETE("/:id/:userId",middleware.AuthMiddleware(),handler.DeleteUserAdress)
	}
}