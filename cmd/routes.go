package main

import (
	"fmt"
	"github.com/SamSweet04/e-commerce_backend.git/controllers"
	"github.com/SamSweet04/e-commerce_backend.git/middlewares"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {
	router := gin.New()

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("auth/model.conf", "auth/policy.csv")
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	//add policy
	//auth.AddPolicies(enforcer)

	router.GET("/roles", controllers.GetRoles)
	auth := router.Group("/auth")
	{
		auth.POST("/token", controllers.GenerateToken)
		auth.POST("/register", controllers.RegisterUser)
	}
	api := router.Group("/api").Use(middlewares.AuthWithJWT())
	{
		api.GET("/ping", middlewares.Authorize("ping", "read", enforcer), controllers.Ping)
	}
	router.Run(":3000")
}
