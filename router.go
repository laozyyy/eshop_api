package main

import (
	"eshop_api/handler"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	userGroup := r.Group("/api/v1/user")
	userGroup.POST("/login", handler.HandleLogin)
	userGroup.POST("/register", handler.HandleRegister)
	userGroup.GET("/get/:uid", handler.HandleGetUser)

	homeGroup := r.Group("/api/v1/home")
	homeGroup.GET("/category/head", handler.HandleGetTags)

	main := r.Group("/api/v1/main")
	{
		main.GET("/get_sku/:sku", handler.HandleGetOneSku)
		main.POST("/mget", handler.HandleMGetSku)
		main.POST("/random", handler.HandleRandom)
		// 在main路由组中添加
		main.POST("/search", handler.HandleSearchGoods)
	}
	cart := r.Group("/api/v1/cart")
	{
		cart.POST("/add", handler.HandleCartAddItem)
		cart.POST("/mget", handler.HandleCartGetList)
		// 在cart路由组中添加
		cart.POST("/update", handler.HandleCartUpdate)
		cart.POST("/delete", handler.HandleCartDelete)
		//cart.POST("/random", handler.HandleRandom)
	}
}
