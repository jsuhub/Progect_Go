package router

import (
	"Progect/controller"
	"Progect/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	log.Println("router初始化")
	home := r.Group("Home")
	{
		home.POST("/login", controller.PassUser)

		home.POST("/register", controller.CreateUser)
	}

	next := r.Group("Next")
	next.Use(middleware.JWTAuthMiddleware())
	{
		next.GET("/get", controller.GetUsers)
	}

	return r
}
