package helper

import (
	"regexp"
)

// IsValidEmail memeriksa apakah string yang diberikan adalah alamat email yang valid.
func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}
