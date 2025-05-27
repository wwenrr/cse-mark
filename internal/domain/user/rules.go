package user

import "regexp"

// IsValidStudentId checks if the provided student ID is valid.
func IsValidStudentId(id string) bool {
	pattern := "^[a-zA-Z0-9]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(id)
}
