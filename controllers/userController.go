package controllers

import (
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"github.com/SamSweet04/e-commerce_backend.git/services"
	"github.com/SamSweet04/e-commerce_backend.git/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	} else {
		user.Password = hashedPassword
	}
	record := database.DB.Create(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func GetUsers(c *gin.Context) {
	users, result := services.GetUsers()
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"users": &users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, result := services.GetUserById(id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")
	result := services.RemoveUser(id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.Status(http.StatusOK)
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	password := c.Query("password")
	result := services.UpdateUser(id, name, password)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.Status(http.StatusOK)
}

func RateItem(c *gin.Context) {
	id := c.Param("id")
	scoreStr := c.Query("rating")
	rating, ratingErr := strconv.ParseFloat(scoreStr, 32)

	if ratingErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Provide valid rating",
		})
	}

	result, err := services.RateItem(id, float32(rating))
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find item",
		})
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}

	c.Status(http.StatusOK)
}

func SaveItem(c *gin.Context) {
	userId := c.Param("userId")
	itemId := c.Query("itemId")

	user, _ := strconv.Atoi(userId)
	item, _ := strconv.Atoi(itemId)

	result := services.SaveItem(user, item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.Status(http.StatusOK)
}

func RemoveSavedItem(c *gin.Context) {
	userId := c.Param("userId")
	itemId := c.Query("itemId")

	user, _ := strconv.Atoi(userId)
	item, _ := strconv.Atoi(itemId)

	result := services.RemoveSavedItem(user, item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.Status(http.StatusOK)
}

func GetSavedItems(c *gin.Context) {
	userId := c.Param("id")
	user, _ := strconv.Atoi(userId)
	items, result := services.GetSavedItem(user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}
