package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID int  `json:"userId" gorm:"not null"`
	User   User `json:"user"`
	ItemID int  `json:"itemId" gorm:"not null"`
	Item   Item `json:"item"`
	Rating int  `json:"rating" gorm:"not null"`
}

// constructor
func NewRating(userId, itemId, rating int) *Rating {
	return &Rating{
		UserID: userId,
		ItemID: itemId,
		Rating: rating,
	}
}
