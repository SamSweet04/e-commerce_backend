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
	//id := c.Param("id")
	userID := c.Param("id")
	user, result := services.GetUserById(userID)
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
	id := c.Param("item_id")
	scoreStr := c.Query("rating")
	itemId, _ := strconv.Atoi(id)
	rating, _ := strconv.Atoi(scoreStr)

	userID, ok := c.Get("id")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}

	if rating < 1 || rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Provide valid rating",
		})
	}
	result, err := services.RateItem(userID.(int), itemId, rating)
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
	itemId := c.Query("itemId")

	item, _ := strconv.Atoi(itemId)
	userID, ok := c.Get("id")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}

	result := services.SaveItem(userID.(int), item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.Status(http.StatusOK)
}

func RemoveSavedItem(c *gin.Context) {
	itemId := c.Query("itemId")

	item, _ := strconv.Atoi(itemId)
	userID, ok := c.Get("userID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}

	result := services.RemoveSavedItem(userID.(int), item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.Status(http.StatusOK)
}

func GetSavedItems(c *gin.Context) {
	userID, ok := c.Get("userID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": ok,
		})
		return
	}
	items, result := services.GetSavedItem(userID.(int))

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

func CommentingItem(c *gin.Context) {
	ratingId := c.Param("ratingId")
	comment := c.Query("comment")
	rating, _ := strconv.Atoi(ratingId)

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
	itemId := c.Param("itemId")
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
	itemId := c.Param("itemId")
	item, _ := strconv.Atoi(itemId)
	userID := c.Param("id")
	user, _ := strconv.Atoi(userID)
	payment, result := services.BuyItem(user, item)
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
