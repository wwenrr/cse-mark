package teleuser

import "regexp"

// IsValidTelegramUsername checks if the provided Telegram username is valid.
func IsValidTelegramUsername(user string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,31}$`
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(user)
}
