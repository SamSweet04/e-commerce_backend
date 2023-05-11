package services

import (
	"fmt"
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"gorm.io/gorm"
	"strconv"
)

func GetItems() ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := database.DB.Find(&items)
	return items, result
}

func GetItem(id string) (models.Item, *gorm.DB) {
	var item models.Item
	result := database.DB.First(&item, id)
	return item, result
}

func AddItem(name, description string, price, userId int) *gorm.DB {
	item := models.NewItem(name, description, price, userId, 0)
	result := database.DB.Create(&item)
	return result
}

func RemoveItem(itemId string) *gorm.DB {
	result := database.DB.Delete(&models.Item{}, itemId)
	return result
}

func UpdateItem(name, itemId string, price int) (*gorm.DB, string) {
	var item models.Item
	result := database.DB.First(&item, itemId)

	if item == (models.Item{}) {
		return result, "No item found"
	}

	item.Name = name
	item.Price = price
	result = database.DB.Save(&item)

	return result, ""
}

func SearchItems(query, orderStr, filter string) ([]models.Item, *gorm.DB) {
	order := ""
	if len(orderStr) > 0 {
		if orderStr[0] == '-' {
			orderStr = orderStr[1:]
			order = fmt.Sprintf("%v desc", orderStr)
		} else {
			order = orderStr
		}
	}
	var items []models.Item
	result := database.DB.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", "%"+query+"%", "%"+query+"%").Order(order).Find(&items)
	return items, result
}
func SearchItemsByDesc(desc, sort string) ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := database.DB.Where("LOWER(description) LIKE ?", "%"+desc+"%").Order(sort).Find(&items)
	return items, result
}
func FilterItemsByPrice(start, end int) ([]models.Item, *gorm.DB) {
	st := strconv.Itoa(start)
	e := strconv.Itoa(end)
	var items []models.Item
	result := database.DB.Where("price BETWEEN ? AND ?", "%"+st+"%", "%"+e+"%").Find(&items)
	return items, result
}
func FilterItemsByRating(start, end int) ([]models.Item, *gorm.DB) {
	st := strconv.Itoa(start)
	e := strconv.Itoa(end)
	var items []models.Item
	result := database.DB.Where("rating BETWEEN ? AND ?", "%"+st+"%", "%"+e+"%").Find(&items)
	return items, result
}
