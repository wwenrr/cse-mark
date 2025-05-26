package validation

import (
	"regexp"
)

// ValidateTelegramUsername validates telegram username format
func ValidateTelegramUsername(user string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,31}$`
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(user)
} 