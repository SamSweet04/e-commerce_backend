package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	RatingID int    `json:"ratingId" gorm:"not null;unique"`
	Rating   Rating `json:"rating"`
	Comment  string `json:"comment" gorm:"not null"`
}

// constructor
func NewComment(ratingId int, comment string) *Comment {
	return &Comment{
		RatingID: ratingId,
		Comment:  comment,
	}
}
