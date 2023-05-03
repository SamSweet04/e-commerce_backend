package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       int     `json:"price" gorm:"not null"`
	Rating      float32 `json:"rating"`
	UserID      int     `json:"userId" gorm:"not null"`
	User        User    `json:"user"`
}

// constructor
func NewItem(name, description string, price, userId int, rating float32) *Item {
	return &Item{
		Name:        name,
		Description: description,
		Price:       price,
		Rating:      rating,
		UserID:      userId,
	}
}
