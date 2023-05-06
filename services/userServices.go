package services

import (
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/handlers"
	"github.com/SamSweet04/e-commerce_backend.git/models"
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
	hashPassword := handlers.HashPassword(password)
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
func RateItem(id string, rating float32) (*gorm.DB, string) {
	var item models.Item
	result := database.DB.First(&item, id)
	if item == (models.Item{}) {
		return result, "No item found"
	}
	item.Rating = (item.Rating + rating) / 2
	result = database.DB.Save(&item)
	return result, ""
}

func SaveItem(userId, itemId int) *gorm.DB {
	savedItem := models.NewSavedItem(userId, itemId)
	result := database.DB.Create(&savedItem)
	return result
}

func RemoveSavedItem(userId, itemId int) *gorm.DB {
	result := database.DB.Where("user_id = ? and item_id = ?", userId, itemId).Delete(models.SavedItem{})
	return result
}

func GetSavedItem(userId int) ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := database.DB.Model(&models.SavedItem{}).Select("items.id, name, description, price, rating").Joins("left join items on savedItem.user_id = ?", userId).Where("items.id = savedItem.item_id").Scan(&items)
	return items, result
}
