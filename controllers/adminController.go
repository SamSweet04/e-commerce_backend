package controllers

import (
	"github.com/SamSweet04/e-commerce_backend.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoles(c *gin.Context) {
	roles, result := services.GetRoles()
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"roles": &roles,
	})
}
