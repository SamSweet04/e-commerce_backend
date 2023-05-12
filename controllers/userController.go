package controllers

import (
	"fmt"
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
		return
	}

	if !utils.IsValidEmail(user.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if !utils.IsValidPassword(user.Password) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Password must have minimum 8 characters and include both letters and numbers"})
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
	userID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := services.UpdateUser(userID, user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.Status(http.StatusOK)
}

func RateItem(c *gin.Context) {
	userID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	user, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	id := c.Param("item_id")
	scoreStr := c.Query("rating")
	itemId, _ := strconv.Atoi(id)
	rating, _ := strconv.Atoi(scoreStr)

	if rating < 1 || rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Provide valid rating",
		})
	}
	result, err := services.RateItem(user, itemId, rating)
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
	userID, ok := c.Get("id")
	itemId := c.Param("item_id")

	item, _ := strconv.Atoi(itemId)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	id, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	result := services.SaveItem(id, item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.Status(http.StatusOK)
}

func RemoveSavedItem(c *gin.Context) {
	userID, ok := c.Get("id")
	itemId := c.Param("item_id")
	item, _ := strconv.Atoi(itemId)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	id := userID.(int)

	result := services.RemoveSavedItem(id, item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.Status(http.StatusOK)
}

func GetSavedItems(c *gin.Context) {
	userID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	id, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	items := services.GetSavedItem(id)

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func CommentingItem(c *gin.Context) {
	ratingId := c.Param("rating_id")
	comment := c.Query("comment")
	rating, _ := strconv.Atoi(ratingId)
	//fmt.Println(ratingId)
	result, err := services.CommentingItem(rating, comment)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
}

func GetComments(c *gin.Context) {
	itemId := c.Param("item_id")
	item, _ := strconv.Atoi(itemId)
	comments, result := services.GetComments(item)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": &comments,
	})
}

func BuyItem(c *gin.Context) {
	userID, ok := c.Get("id")
	itemId := c.Param("itemId")
	item, _ := strconv.Atoi(itemId)

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	id, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	payment, result := services.BuyItem(id, item)
	if result != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Payment": &payment,
	})
}
