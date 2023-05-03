package models

import "gorm.io/gorm"

type SavedItem struct {
	gorm.Model
	UserID int  `json:"userId" gorm:"not null"`
	User   User `json:"user"`
	ItemID int  `json:"itemId" gorm:"not null"`
	Item   Item `json:"item"`
}

// constructor
func NewSavedItem(userId, itemId int) *SavedItem {
	return &SavedItem{
		UserID: userId,
		ItemID: itemId,
	}
}
