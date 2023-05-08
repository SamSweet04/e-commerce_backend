package services

import (
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"gorm.io/gorm"
)

func GetRoles() ([]models.Role, *gorm.DB) {
	var roles []models.Role
	result := database.DB.Find(&roles)
	return roles, result
}
