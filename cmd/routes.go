package main

import (
	"fmt"
	"github.com/SamSweet04/e-commerce_backend.git/controllers"
	"github.com/SamSweet04/e-commerce_backend.git/middlewares"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("auth/model.conf", "auth/policy.csv")
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	// Public routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// Authenticated routes
	authenticated := router.Group("/auth")
	authenticated.Use(middlewares.AuthWithJWT())
	{
		authenticated.GET("/", controllers.WelcomePage)
		authenticated.GET("/getusers", middlewares.Authorize("user", "read", enforcer), controllers.GetUsers)

		// User routes
		userRoutes := authenticated.Group("/user")
		{
			userRoutes.GET("/", controllers.GetUsers)
			userRoutes.GET("/:id", controllers.GetUserById)
			userRoutes.PUT("/", controllers.UpdateUser)
			userRoutes.DELETE("/:id", controllers.RemoveUser)
			userRoutes.POST("/save/:item_id", controllers.SaveItem)
			userRoutes.DELETE("/savedRemove/:item_id", controllers.RemoveSavedItem)
			userRoutes.GET("/save/", controllers.GetSavedItems)
			userRoutes.PUT("/rate/:item_id", controllers.RateItem)
			userRoutes.POST("/comment/:rating_id", controllers.CommentingItem)
			userRoutes.GET("/comment/:item_id", controllers.GetComments)
			userRoutes.PUT("/:id/buy/:itemId", controllers.BuyItem)
		}
		// Define routes for managing items
		items := router.Group("/items")
		{
			items.GET("/", controllers.GetItems)
			items.GET("/:id", controllers.GetItem)
			items.POST("/", controllers.AddItem)
			items.PUT("/:id", controllers.UpdateItem)
			items.DELETE("/:id", controllers.RemoveItem)
			items.GET("/search", controllers.SearchItems)
		}
	}
	router.Run(":3000")
}
