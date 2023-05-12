package utils

import (
	"regexp"
	"strings"
)

// Define validation rules for email and password
var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// Function to validate email format
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	// Password must have minimum 8 characters and include both letters and numbers
	re := regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`)
	return re.MatchString(password) && strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}
