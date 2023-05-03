package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID   int    `json:"userId" gorm:"not null"`
	User     User   `json:"user"`
	ItemID   int    `json:"itemId" gorm:"not null"`
	Item     Item   `json:"item"`
	RatingID int    `json:"ratingId" gorm:"not null;unique"`
	Rating   Rating `json:"rating"`
	Comment  string `json:"comment" gorm:"not null"`
}

// constructor
func NewComment(userId, itemId, ratingId int, comment string) *Comment {
	return &Comment{
		UserID:   userId,
		ItemID:   itemId,
		RatingID: ratingId,
		Comment:  comment,
	}
}
