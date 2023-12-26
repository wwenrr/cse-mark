package validation

import (
	"regexp"
)

func ValidateStudentId(id string) bool {
	pattern := "^[a-zA-Z0-9]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(id)
}

func ValidateCourseId(id string) bool {
	pattern := "^[a-zA-Z][a-zA-Z0-9-]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(id)
}

func ValidateTelegramUsername(user string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,31}$`
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(user)
}
