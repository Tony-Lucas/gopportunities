package router

import (
	"github.com/Tony-Lucas/gopportunities/handler"
	"github.com/Tony-Lucas/gopportunities/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeProductRoutes(r *gin.Engine) {
	v1 := r.Group("/product")
	{
		v1.GET("/:id", handler.GetProduct)
		v1.GET("/limit/:offset", handler.GetProducts)
		v1.POST("/", middleware.AdminMiddleware(), handler.PostProduct)
		v1.PUT("/", middleware.AdminMiddleware(), handler.UpdateProduct)
		v1.DELETE("/:id", middleware.AdminMiddleware(), handler.DeleteProduct)
	}

}
