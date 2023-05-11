package database

import (
	"fmt"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Almaty",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.SavedItem{})
	db.AutoMigrate(&models.Rating{})
	db.AutoMigrate(&models.Comment{})

	DB = db
}
