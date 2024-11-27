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
}
