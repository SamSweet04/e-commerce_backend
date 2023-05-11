package controllers

import (
	"github.com/SamSweet04/e-commerce_backend.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WelcomePage(c *gin.Context) {
	userID, ok := c.Get("id")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	user, result := services.GetUserById(userID.(string))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + user.Username + "!"})
}
