package helper

import "strings"

func ValidateData(userFirstName string, userLastName string, userEmail string) (bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	
	return isValidName, isValidEmail
}