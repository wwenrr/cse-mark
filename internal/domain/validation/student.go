package validation

import (
	"regexp"
)

// ValidateStudentId validates student ID format
func ValidateStudentId(id string) bool {
	pattern := "^[a-zA-Z0-9]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(id)
}
