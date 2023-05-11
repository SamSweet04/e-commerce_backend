package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// validation checks
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func CheckPassword(providedPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(providedPassword))
	return err == nil
}
