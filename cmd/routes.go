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

		// User routes
		userRoutes := authenticated.Group("/user")
		{
			userRoutes.GET("/", middlewares.Authorize("user", "read", enforcer), controllers.GetUsers)
			userRoutes.GET("/:id", middlewares.Authorize("user", "read", enforcer), controllers.GetUserById)
			userRoutes.PUT("/", middlewares.Authorize("user", "write", enforcer), controllers.UpdateUser)
			userRoutes.DELETE("/:id", middlewares.Authorize("user", "write", enforcer), controllers.RemoveUser)
			userRoutes.POST("/save/:item_id", controllers.SaveItem)
			userRoutes.DELETE("/savedRemove/:item_id", controllers.RemoveSavedItem)
			userRoutes.GET("/save/", controllers.GetSavedItems)
			userRoutes.PUT("/rate/:item_id", controllers.RateItem)
			userRoutes.POST("/comment/:rating_id", controllers.CommentingItem)
			userRoutes.GET("/comment/:item_id", controllers.GetComments)
			userRoutes.PUT("/buy/:itemId", controllers.BuyItem)
		}
		// Define routes for managing items
		items := authenticated.Group("/items")
		{
			items.GET("/", controllers.GetItems)
			items.GET("/:id", controllers.GetItem)
			items.POST("/", middlewares.Authorize("item", "write", enforcer), controllers.AddItem)
			items.PUT("/:id", middlewares.Authorize("item", "write", enforcer), controllers.UpdateItem)
			items.DELETE("/:id", middlewares.Authorize("item", "write", enforcer), controllers.RemoveItem)
			items.GET("/search", controllers.SearchItems)
		}
	}
	router.Run(":3000")
}
