package main

import (
	"eshop_api/handler"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	group := r.Group("/api/v1/user")
	group.POST("/login", handler.HandleLogin)
	group.POST("/register", handler.HandleRegister)
	group.GET("/get/:uid", handler.HandleGetUser)
}
