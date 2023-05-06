package main

import (
	"github.com/SamSweet04/e-commerce_backend.git/controllers"
	"github.com/SamSweet04/e-commerce_backend.git/middlewares"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)

		}
	}
	return router
}
