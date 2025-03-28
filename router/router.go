package router

import (
	"Progect/controller"
	"Progect/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
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
