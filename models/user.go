package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Telegram  string `json:"telegram"`
	RoleID    int    `json:"roleId" gorm:"not null"`
	Role      Role   `json:"role"`
}

// constructor
func NewUser(username, password, email string, roleId int) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
		RoleID:   roleId,
	}
}
