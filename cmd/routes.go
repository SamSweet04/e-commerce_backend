package main

import (
	"fmt"
	"github.com/SamSweet04/e-commerce_backend.git/auth"
	"github.com/SamSweet04/e-commerce_backend.git/controllers"
	"github.com/SamSweet04/e-commerce_backend.git/middlewares"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	enforcer, err := auth.NewEnforcer()
	if err != nil {
		panic(err)
	}
	router.Use(authz.NewAuthorizer(enforcer))

	router.GET("/roles", controllers.GetRoles)
	router.GET("/data1/read", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You can read data1"})
	})
	router.GET("/data1/write", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You can write data1"})
	})
	api := router.Group("/auth")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	router.Run(":3000")
}
