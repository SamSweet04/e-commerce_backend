package models

import (
	"golang.org/x/crypto/bcrypt"
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
} //not used

// validation checks
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
