package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Telegram  string `json:"telegram"`
	Role      string `json:"role"`
	Balance   int    `json:"balance" gorm:"not null"`
}

// constructor
func NewUser(username, password, email, role string) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
		Role:     role,
		Balance:  0,
	}
}
