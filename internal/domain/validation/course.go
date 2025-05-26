package validation

import (
	"regexp"
)

// ValidateCourseId validates course ID format
func ValidateCourseId(id string) bool {
	pattern := "^[a-zA-Z][a-zA-Z0-9-]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(id)
}
