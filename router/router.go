package router

import (
	"eshop_api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1/main")
	{
		v1.GET("/get_sku/:sku", handler.HandleGetOneSku)
		v1.POST("/mget", handler.HandleMGetSku)
	}
}
