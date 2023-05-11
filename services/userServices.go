package services

import (
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"github.com/SamSweet04/e-commerce_backend.git/utils"
	"gorm.io/gorm"
)

func GetUsers() ([]models.User, *gorm.DB) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result
}

func GetUserById(id string) (models.User, *gorm.DB) {
	var user models.User
	result := database.DB.First(&user, id)
	return user, result
}

func RemoveUser(id string) *gorm.DB {
	result := database.DB.Delete(&models.User{}, id)
	return result
}

func UpdateUser(id, name, password string) *gorm.DB {
	hashPassword, _ := utils.HashPassword(password)
	var user models.User
	result := database.DB.First(&user, id)
	if user == (models.User{}) {
		return result
	}
	user.Username = name
	user.Password = hashPassword
	result = database.DB.Save(&user)
	return result
}
func RateItem(userID, itemId, rating int) (*gorm.DB, string) {
	var item models.Item
	result := database.DB.First(&item, itemId)
	if item == (models.Item{}) {
		return result, "No item found"
	}
	newRating := models.NewRating(userID, itemId, rating)
	res := database.DB.Create(&newRating)
	item.Rating = CalculateRating(itemId)
	result = database.DB.Save(&item)
	return res, ""
}

func CalculateRating(itemId int) float32 {
	var rating []models.Rating
	database.DB.Find("itemId = ?", itemId, &rating)
	sum := 0
	for i := 0; i < len(rating); i++ {
		sum += rating[i].Rating
	}
	return float32(sum / len(rating))

}

func SaveItem(userId, itemId int) *gorm.DB {
	savedItem := models.NewSavedItem(userId, itemId)
	result := database.DB.Create(&savedItem)
	return result
}

func RemoveSavedItem(userId, itemId int) *gorm.DB {
	result := database.DB.Where("userID = ? and itemID = ?", userId, itemId).Delete(models.SavedItem{})
	return result
}

func GetSavedItem(userId int) ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := database.DB.Model(&models.SavedItem{}).Select("items.id, name, description, price, rating").Joins("left join items on savedItem.user_id = ?", userId).Where("items.id = savedItem.item_id").Scan(&items)
	return items, result
}

func CommentingItem(ratingId int, comment string) (*gorm.DB, string) {
	var rating models.Rating
	result := database.DB.First(&rating, ratingId)
	if rating == (models.Rating{}) {
		return result, "No item found"
	}
	newComment := models.NewComment(ratingId, comment)
	res := database.DB.Create(&newComment)
	return res, ""
}

func GetComments(itemId int) ([]models.Comment, *gorm.DB) {
	var comments []models.Comment
	result := database.DB.Find(&comments, itemId)
	return comments, result
}

func BuyItem(userID, itemId int) (*gorm.DB, bool) {
	var item models.Item
	result := database.DB.First(&item, itemId)
	if item == (models.Item{}) {
		return result, false
	}
	var user models.User
	result2 := database.DB.First(&user, userID)
	if user == (models.User{}) {
		return result2, false
	}
	var seller models.User
	result3 := database.DB.First(&seller, item.UserID)
	if seller == (models.User{}) {
		return result3, false
	}
	if user.Balance < item.Price {
		return nil, false
	} else {
		user.Balance -= item.Price
		seller.Balance += item.Price
	}
	result = database.DB.Save(&user)
	result3 = database.DB.Save(&seller)
	return result, true
}
