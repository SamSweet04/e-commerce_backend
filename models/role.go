package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role string `json:"role" gorm:"not null"`
}

// constructor
func NewRole(role string) *Role {
	return &Role{
		Role: role,
	}
}
