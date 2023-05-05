package main

import (
	"github.com/SamSweet04/e-commerce_backend.git/controllers"
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/middlewares"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()
	role := models.Role{Role: "customer"}
	database.DB.Create(&role)
	//var user = models.User{}
	//database.DB.First(&user)
	//fmt.Println(user)
	router := initRouter()
	router.Run(":8080")
}

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
