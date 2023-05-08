package database

import (
	"errors"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"gorm.io/gorm"
)

func HandleRoles(db *gorm.DB) *gorm.DB {
	adminRole := models.NewRole("admin")
	//db.Exec("TRUNCATE TABLE roles RESTART IDENTITY CASCADE;")
	result := db.First(&adminRole)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&adminRole)
	}

	return result
}
